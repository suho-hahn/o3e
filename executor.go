package o3e

import (
    "sync/atomic"
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
        make(chan bool),
        0,
    }

}

func (e *Executor) Start() {

    for i := range e.channels {
        go e.handleChannel(i)
    }
}

func (e *Executor) Stop() {
    e.stopCh <- true
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
        case wrap := <- ch: e.executeTaskWrapper(wrap)
        case <- e.stopCh: break LOOP
        }

    }

    //e.Stop()

}

func (e *Executor) executeTaskWrapper(wrap *taskWrapper) {
    result, err := wrap.execute()
    execResultHandlers[result](e, wrap, err)
}

var execResultHandlers = []func(*Executor, *taskWrapper, error) {
    handleResultSuccess,
    handleResultWait,
    handleResultError,
}

func handleResultSuccess(e *Executor, wrap *taskWrapper, _ error) {
    // TODO not implemented yet
}

func handleResultWait(e *Executor, wrap *taskWrapper, _ error) {
    // TODO not implemented yet
}

func handleResultError(e *Executor, wrap *taskWrapper, _ error) {
    // TODO not implemented yet

    // TODO Pass Error
    //e.Stop()
}