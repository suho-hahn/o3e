package o3e

import (
    "testing"
    "errors"
)

type PrintTask struct {
    Deps      map[int]Void
    Str       string
    ExecCount int
    Test      *testing.T
    isPanic   bool
    isError   bool
}

func (t *PrintTask) DepFactors() map[int]Void {
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



// copied from https://gist.github.com/samalba/6059502#gistcomment-2331327
func assertEqual(t *testing.T, a interface{}, b interface{}) {
    t.Log("assertEqual", a, b)
    if a != b {
        t.Fatalf("%v != %v", a, b)
        t.FailNow()
    }
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}) {
    t.Log("assertEqual", a, b)
    if a == b {
        t.Fatalf("%v == %v", a, b)
        t.FailNow()
    }
}