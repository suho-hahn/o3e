package o3e

import (
    "sync/atomic"
    "errors"
    "fmt"
)

type wrapperResult uint8
type EmptyType struct{}
var Empty = EmptyType{}

const (
    wrapperSuccess wrapperResult = iota
    wrapperWait
    wrapperError
)

type Task interface {
    DepFactors() map[int]EmptyType // memoization may improve performance.
    Execute() error
    // TODO Commit
    // TODO Rollback()
}

type taskWrapper struct {
    Task
    blockCount int32
    successCh chan bool
}

func newTaskWrapper(t Task) *taskWrapper {

    deps := t.DepFactors()

    result := &taskWrapper{
        t,
        int32(len(deps)),
        make(chan bool, 1),
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
        err = w.Execute()
        if err != nil {
            return wrapperError, err
        } else {
            return wrapperSuccess, nil
        }
    } else {
        return wrapperWait, nil
    }
}
