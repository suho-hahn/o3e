package o3e

import (
    "sync/atomic"
)



type Task interface {
    DepFactors() map[int]bool
    Execute()
}

type taskWrappper struct {
    Task
    deps   map[int]bool //memoizatiojn
    waitCount int32
}

func newTaskWrapper(t Task) *taskWrappper {

    deps := t.DepFactors()

    result := &taskWrappper{
        t,
        deps,
        int32(len(deps)),
    }

    return result

}

func (w *taskWrappper) execute() {
    if atomic.AddInt32(&w.waitCount, -1) == 0 {
        w.Execute()
    }
}
