package o3e

import (

)

type Executor struct {
    queueList []chan *taskWrap
    nextFetchId int64
    nextCommitId int64

}

func NewExecutor(queueListSize, queueSize int) *Executor {

    queueList := make([]chan *taskWrap, queueListSize)
    for i := range queueList {
        queueList[i] = make(chan *taskWrap, queueSize)
    }

    return &Executor {
        queueList,
        0,
        0,
    }

}

func (e *Executor) Start() {
    // TODO
}

func (e *Executor) Stop() {
    // TODO
}

// Thread **UNSAFE**
func (e *Executor) AddTask(t Task) {

    wrap := newTaskWrap(t)


    for depFactor := range wrap.deps {
        e.queueList[depFactor % len(e.queueList)] <- wrap
    }

}

func (e *Executor) handleQueueAsync(queueIndex int) {

    ch := e.queueList[queueIndex]

    for {
        wrap := <- ch
        wrap.execute()
    }

}
