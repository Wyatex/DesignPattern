package singleton

import (
    "fmt"
    "testing"
)

func TestSyncOnceSingleton(t *testing.T) {
    for i := 0; i < 30; i++ {
        go getOnceInstance()
    }

    fmt.Scanln()
}
