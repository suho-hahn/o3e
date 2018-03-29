package o3e

import (
    "sync/atomic"
)



type Task interface {
    DepFactors() map[int]bool
    Execute()
}

type taskWrap struct {
    Task
    deps   map[int]bool //memoizatiojn
    waitCount int32
    doneCh chan bool
    stopCh chan bool
}

func newTaskWrap(t Task, stopCh chan bool) *taskWrap {

    deps := t.DepFactors()

    result := &taskWrap{
        t,
        deps,
        int32(len(deps)),
        make(chan bool, len(deps)),
        stopCh,
    }

    return result

}

func (w *taskWrap) execute() {
    if atomic.AddInt32(&w.waitCount, -1) != 0 {
        select {
        case <-w.stopCh:
        case <-w.doneCh:
        }
    } else {
        defer w.done()
        w.Execute()
    }
}

func (w *taskWrap) done() {
    for i:=0; i<len(w.deps) - 1; i++ {
        w.doneCh <- true
    }
}