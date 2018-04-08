package o3e

import (
    "sync/atomic"
    "log"
)

type Executor struct {
    channels []chan *taskWrapper
    stopCh   chan bool
    numGoroutine int32
}

func NewExecutor(numOfChannels, channelCapacity int) *Executor {

    channels := make([]chan *taskWrapper, numOfChannels)
    for i := range channels {
        channels[i] = make(chan *taskWrapper, channelCapacity)
    }

    return &Executor {
        channels,
        make(chan bool, numOfChannels),
        0,
    }

}

func (e *Executor) Start() {

    for i := range e.channels {
        go e.handleChannel(i)
    }
}

func (e *Executor) Stop() {
    LOOP: for {
        select {
        case e.stopCh <- true:
        default:
            break LOOP
        }
    }

}

// Thread **UNSAFE**
func (e *Executor) AddTask(t Task) {

    wrap := newTaskWrapper(t)

    deps := wrap.DepFactors()
    for depFactor := range deps {
        e.channels[depFactor % len(e.channels)] <- wrap
    }

}

func (e *Executor) handleChannel(chIdx int) {

    atomic.AddInt32(&e.numGoroutine, 1)
    defer atomic.AddInt32(&e.numGoroutine, -1)

    ch := e.channels[chIdx]

    LOOP: for {

        // check stop
        select {
        case <- e.stopCh: break LOOP
        default: // DO NOTHING
        }

        // execute or stop
        select {
        case <- e.stopCh: break LOOP
        case wrap := <- ch:

            result, err := wrap.execute()
            if result == wrapperWait {
                isSuccess := <- wrap.successCh // wait result
                wrap.successCh <- isSuccess
                if ! isSuccess {
                    break LOOP // stop
                }
            } else if result == wrapperSuccess {
                wrap.successCh <- true
            } else { // wrapperError
                wrap.successCh <- false
                log.Print(err) // FIXME
                break LOOP
            }

        }

    }

    e.Stop()

}