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

type taskWrapper struct {
    Task
    blockCount int32
}

func newTaskWrapper(t Task) *taskWrapper {

    deps := t.DepFactors()

    result := &taskWrapper{
        t,
        int32(len(deps)),
    }

    return result

}

func (w *taskWrapper) execute() (result wrapperResult, err error) {

    defer func() {
        rec := recover()
        if rec != nil {
            result = wrapperError
            err = errors.New(fmt.Sprint(rec))
            return
        }
    }()

    if atomic.AddInt32(&w.blockCount, -1) == 0 {
        w.Execute()
        return wrapperSuccess, nil
    } else {
        return wrapperWait, nil
    }
}
