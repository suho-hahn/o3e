package o3e

import (
    "testing"
    "time"
    "runtime"
)

func TestTaskWrap_NormalExecution(t *testing.T) {

    t.Log(t.Name())

    task := &PrintTask{
        map[int]bool{1:false, 2:false, 3:false},
        "PrintTask is working",
        0,
        t,
    }

    stopCh := make(chan bool)

    w := newTaskWrap(task, stopCh)
    exNumGo, curNumGo := 0, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(3))

    t.Log("exec")
    go w.execute()
    time.Sleep(time.Second)
    exNumGo, curNumGo = curNumGo, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(2))
    assertEqual(t, exNumGo+ 1, curNumGo)

    t.Log("exec")
    go w.execute()
    time.Sleep(time.Second)
    exNumGo, curNumGo = curNumGo, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(1))
    assertEqual(t, exNumGo+ 1, curNumGo)

    t.Log("exec")
    go w.execute()
    time.Sleep(time.Second)
    exNumGo, curNumGo = curNumGo, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 1)
    assertEqual(t, w.waitCount, int32(0))
    assertEqual(t, exNumGo - 2, curNumGo)

}

func TestTaskWrap_Interrupt(t *testing.T) {

    t.Log(t.Name())

    task := &PrintTask{
        map[int]bool{1:false, 2:false, 3:false},
        "PrintTask is working",
        0,
        t,
    }

    stopCh := make(chan bool)

    w := newTaskWrap(task, stopCh)
    exNumGo, curNumGo := 0, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(3))

    t.Log("exec")
    go w.execute()
    time.Sleep(time.Second)
    exNumGo, curNumGo = curNumGo, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(2))
    assertEqual(t, exNumGo+ 1, curNumGo)

    t.Log("exec")
    go w.execute()
    time.Sleep(time.Second)
    exNumGo, curNumGo = curNumGo, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(1))
    assertEqual(t, exNumGo+ 1, curNumGo)

    t.Log("stop")
    stopCh <- true
    time.Sleep(time.Second)
    exNumGo, curNumGo = curNumGo, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(1))
    assertEqual(t, exNumGo - 1, curNumGo)

    t.Log("stop")
    stopCh <- true
    time.Sleep(time.Second)
    exNumGo, curNumGo = curNumGo, runtime.NumGoroutine()

    assertEqual(t, task.ExecCount, 0)
    assertEqual(t, w.waitCount, int32(1))
    assertEqual(t, exNumGo - 1, curNumGo)

}