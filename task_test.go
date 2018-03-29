package o3e

import (
    "testing"
    "time"
)

func TestTaskWrap_NormalExecution(t *testing.T) {

    t.Log(t.Name())

    task := &PrintTask{
        map[int]bool{1:false, 2:false, 3:false},
        "PrintTask is working",
        0,
        t,
    }

    w := newTaskWrapper(task)

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(3))

    t.Log("exec")
    go w.execute()
    time.Sleep(time.Second)

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(2))

    t.Log("exec")
    go w.execute()
    time.Sleep(time.Second)

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(1))

    t.Log("exec")
    go w.execute()
    time.Sleep(time.Second)

    assertEqual(t, task.ExecCount, 1)
    assertEqual(t, w.waitCount, int32(0))

}
