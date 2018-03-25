package o3e

import (
    "sync"
    "sync/atomic"
)



type Task interface {
    DepFactors() map[int]bool
    Execute()
}

type taskWrap struct {
    Task
    deps      map[int]bool //memoizatiojn
    waitCount int32
    wait      sync.WaitGroup
}

func newTaskWrap(t Task) *taskWrap {

    deps := t.DepFactors()

    result := &taskWrap{
        t,
        deps,
        int32(len(deps)),
        sync.WaitGroup{},
    }

    result.wait.Add(1)

    return result

}

func (w *taskWrap) execute() {
    if atomic.AddInt32(&w.waitCount, -1) != 0 {
        w.wait.Wait()
    } else {
        defer w.wait.Done()
        w.Execute() // TODO error handling
    }
}
