package o3e

import (
    "testing"
    "runtime"
)

// copied from https://gist.github.com/samalba/6059502#gistcomment-2331327
func assertEqual(t *testing.T, a interface{}, b interface{}) {

    _, file, line, _ := runtime.Caller(1)

    t.Log("assertEqual", file, line, a, b)
    if a != b {
        t.Fatalf("%v != %v", a, b)
        t.FailNow()
    }
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}) {
    _, file, line, _ := runtime.Caller(1)
    t.Log("assertEqual", file, line, a, b)
    if a == b {
        t.Fatalf("%v == %v", a, b)
        t.FailNow()
    }
}
