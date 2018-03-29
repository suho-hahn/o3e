package o3e

import (
    "sync/atomic"
)

type Executor struct {
    channels []chan *taskWrappper
    stopCh   chan bool
    numGoroutine int32
}

func NewExecutor(numOfChannels, channelCapacity int) *Executor {

    channels := make([]chan *taskWrappper, numOfChannels)
    for i := range channels {
        channels[i] = make(chan *taskWrappper, channelCapacity)
    }

    return &Executor {
        channels,
        make(chan bool),
        0,
    }

}

func (e *Executor) Start() {

    for i := range e.channels {
        go e.handleQueueAsync(i)
        atomic.AddInt32(&e.numGoroutine, 1)
    }
}

func (e *Executor) Stop() {
    e.stopCh <- true
}

// Thread **UNSAFE**
func (e *Executor) AddTask(t Task) {

    wrap := newTaskWrapper(t, e.stopCh)

    for depFactor := range wrap.deps {
        e.channels[depFactor % len(e.channels)] <- wrap
    }

}

func (e *Executor) handleQueueAsync(chIdx int) {

    ch := e.channels[chIdx]

    for {

        select {
        case <- e.stopCh:
            goto STOP
        default:
        }

        select {
        case wrap := <- ch:
            wrap.execute()
        case <- e.stopCh:
            e.stopCh <- true
            goto STOP
        }

    }

    STOP:
    atomic.AddInt32(&e.numGoroutine, -1)

}
