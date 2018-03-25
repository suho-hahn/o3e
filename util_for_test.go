package o3e

import (
    "fmt"
    "testing"
)

type PrintTask struct {
    Deps map[int]bool
    Str string
    ExecCount int
}

func (t *PrintTask) DepFactors() map[int]bool {
    return t.Deps
}

func (t *PrintTask) Execute() {
    fmt.Println(t.Str)
    t.ExecCount ++
}

// copied from https://gist.github.com/samalba/6059502#gistcomment-2331327
func assertEqual(t *testing.T, a interface{}, b interface{}) {
    if a != b {
        t.Fatalf("%s != %s", a, b)
    }
}