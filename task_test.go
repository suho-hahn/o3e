package o3e

import (
    "testing"
    "time"
    "runtime"
)

func TestNewTaskWrap(t *testing.T) {

    task := &PrintTask{
        map[int]bool{1:false, 2:false},
        "PrintTask is working",
        0,
    }

    w := newTaskWrap(task)
    numGo1, numGo2 := 0, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(2))

    go w.execute()
    time.Sleep(time.Second)
    numGo1, numGo2 = numGo2, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(1))
    assertEqual(t, numGo1 + 1, numGo2)

    go w.execute()
    time.Sleep(time.Second)
    numGo1, numGo2 = numGo2, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 1)
    assertEqual(t, w.waitCount, int32(0))
    assertEqual(t, numGo1 - 1, numGo2)

}