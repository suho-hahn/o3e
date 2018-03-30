package o3e

import (
    "sync/atomic"
    "errors"
    "fmt"
)

type wrapperResult uint8
type Void bool

const (
    wrapperSuccess wrapperResult = iota
    wrapperWait
    wrapperError
)

type Task interface {
    DepFactors() map[int]Void // memoization may improve performance.
    Execute()
}

type taskWrappper struct {
    Task
    waitCount int32
}

func newTaskWrapper(t Task) *taskWrappper {

    deps := t.DepFactors()

    result := &taskWrappper{
        t,
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
