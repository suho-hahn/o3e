package o3e

import (
    "time"
    "sync/atomic"
    "testing"
    "runtime"
    "fmt"
)

// TODO start & stop

type SleepTask struct {
    ExecCounter *int32
    SleepDuration time.Duration
    Dependency int
}

func (t *SleepTask) DepFactors() map[int]EmptyType {
    deps := make(map[int]EmptyType)
    deps[t.Dependency] = Empty
    return deps
}

func (t *SleepTask) Execute() error {
    fmt.Println("start sleep task", t.Dependency)
    time.Sleep(t.SleepDuration)
    atomic.AddInt32(t.ExecCounter, 1)
    fmt.Println("end sleep task", t.Dependency)
    return nil
}

func TestExecutor_start_stop(t *testing.T){

    go1 := runtime.NumGoroutine()

    var execCount int32
    exec := NewExecutor(8, 8)
    exec.Start()
    time.Sleep(time.Second)
    go2 := runtime.NumGoroutine()
    assertEqual(t, go1 + 8, go2)

    exec.AddTask(&SleepTask{&execCount, 5 * time.Second, 1})
    exec.AddTask(&SleepTask{&execCount, 5 * time.Second, 2})
    exec.AddTask(&SleepTask{&execCount, 5 * time.Second, 3})
    exec.AddTask(&SleepTask{&execCount, 5 * time.Second, 1})
    exec.AddTask(&SleepTask{&execCount, 5 * time.Second, 2})
    exec.AddTask(&SleepTask{&execCount, 5 * time.Second, 3})

    time.Sleep(6 * time.Second)
    assertEqual(t, int32(3), execCount)

    exec.Stop()
    time.Sleep(time.Second)
    go3 := runtime.NumGoroutine()
    assertEqual(t, go1 + 3, go3)
    assertEqual(t, int32(3), execCount)

    time.Sleep(5 * time.Second)
    go4 := runtime.NumGoroutine()
    assertEqual(t, go1, go4)
    assertEqual(t, int32(6), execCount)

}

// TODO dependency wait
// TODO exec error
