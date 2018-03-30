package o3e

import (
    "testing"
)

func TestTaskWrap_NormalExecution(t *testing.T) {

    t.Log(t.Name())

    task := &PrintTask{
        map[int]Void{1:false, 2:false, 3:false},
        "PrintTask is working",
        0,
        t,
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

func TestTaskWrap_Error(t *testing.T) {

    t.Log(t.Name())

    task := &PrintTask{
        map[int]Void{1:false, 2:false, 3:false},
        "PrintTask is working",
        0,
        t,
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
    assertNotEqual(t, err, nil)
    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.blockCount, int32(0))

}

func TestWrapperResultValue(t *testing.T) {
    assertEqual(t, wrapperSuccess + 1, wrapperWait)
    assertEqual(t, wrapperWait + 1, wrapperError)
}