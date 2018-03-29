package o3e

import (
    "testing"
)

type PrintTask struct {
    Deps      map[int]bool
    Str       string
    ExecCount int
    Test      *testing.T
}

func (t *PrintTask) DepFactors() map[int]bool {
    return t.Deps
}

func (t *PrintTask) Execute() {
    t.Test.Log(t.Str)
    t.ExecCount ++
}

// copied from https://gist.github.com/samalba/6059502#gistcomment-2331327
func assertEqual(t *testing.T, a interface{}, b interface{}) {
    t.Log("assertEqual", a, b)
    if a != b {
        t.Fatalf("%v != %v", a, b)
    }
}