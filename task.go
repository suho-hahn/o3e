package o3e

import (
    "sync/atomic"
    "errors"
    "fmt"
)

type wrapperResult uint8

const (
    wrapperSuccess wrapperResult = iota
    wrapperWait
    wrapperError
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

func (w *taskWrappper) execute() (result wrapperResult, err error) {

    defer func() {
        rec := recover()
        if rec != nil {
            result = wrapperError
            err = errors.New(fmt.Sprint(rec))
            return
        }
    }()

    if atomic.AddInt32(&w.waitCount, -1) == 0 {
        w.Execute()
        return wrapperSuccess, nil
    } else {
        return wrapperWait, nil
    }
}
