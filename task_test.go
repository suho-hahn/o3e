package o3e

import (
    "testing"
    "fmt"
    "errors"
)

type PrintTask struct {
    Deps      map[int]EmptyType
    Str       string
    ExecCount int
    Test      *testing.T
    isPanic   bool
    isError   bool
}

func (t *PrintTask) DepFactors() map[int]EmptyType {
    return t.Deps
}

func (t *PrintTask) Execute() error {

    if t.isPanic {
        panic("make panic")
    }

    if t.isError {
        return errors.New("make error")
    }

    t.Test.Log(t.Str)
    t.ExecCount ++

    return nil
}

func TestTaskWrap_NormalExecution(t *testing.T) {

    t.Log(t.Name())

    task := &PrintTask{
        map[int]EmptyType{1:Empty, 2:Empty, 3:Empty},
        "PrintTask is working",
        0,
        t,
        false,
        false,
    }

    w := newTaskWrapper(task)
    var result wrapperResult
    var err error

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(3))

    t.Log("exec")
    result, err = w.execute()
    assertEqual(t, result, wrapperWait)
    assertEqual(t, err, nil)
    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(2))

    t.Log("exec")
    result, err = w.execute()
    assertEqual(t, result, wrapperWait)
    assertEqual(t, err, nil)
    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(1))

    t.Log("exec")
    result, err = w.execute()
    assertEqual(t, result, wrapperSuccess)
    assertEqual(t, err, nil)
    assertEqual(t, task.ExecCount, 1)
    assertEqual(t, w.blockCount, int32(0))

}

func TestTaskWrap_Panic(t *testing.T) {

    t.Log(t.Name())

    task := &PrintTask{
        map[int]EmptyType{1:Empty, 2:Empty, 3:Empty},
        "PrintTask is working",
        0,
        t,
        true,
        false,
    }

    w := newTaskWrapper(task)
    var result wrapperResult
    var err error

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(3))

    t.Log("exec")
    result, err = w.execute()
    assertEqual(t, result, wrapperWait)
    assertEqual(t, err, nil)
    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(2))

    t.Log("exec")
    result, err = w.execute()
    assertEqual(t, result, wrapperWait)
    assertEqual(t, err, nil)
    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(1))

    t.Log("exec")
    result, err = w.execute()
    assertEqual(t, result, wrapperError)
    assertEqual(t, fmt.Sprint(err), "make panic")
    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(0))

}

func TestTaskWrap_Error(t *testing.T) {

    t.Log(t.Name())

    task := &PrintTask{
        map[int]EmptyType{1:Empty, 2:Empty, 3:Empty},
        "PrintTask is working",
        0,
        t,
        false,
        true,
    }

    w := newTaskWrapper(task)
    var result wrapperResult
    var err error

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(3))

    t.Log("exec")
    result, err = w.execute()
    assertEqual(t, result, wrapperWait)
    assertEqual(t, err, nil)
    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(2))

    t.Log("exec")
    result, err = w.execute()
    assertEqual(t, result, wrapperWait)
    assertEqual(t, err, nil)
    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(1))

    t.Log("exec")
    result, err = w.execute()
    assertEqual(t, result, wrapperError)
    assertEqual(t, fmt.Sprint(err),  "make error")
    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(0))

}

func TestWrapperResultValue(t *testing.T) {
    assertEqual(t, wrapperSuccess + 1, wrapperWait)
    assertEqual(t, wrapperWait + 1, wrapperError)
}