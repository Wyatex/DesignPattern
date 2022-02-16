package singleton

import (
    "fmt"
    "sync"
)

var once sync.Once

type singleSyncOnce struct {
}

var singleOnceInstance *singleSyncOnce

func getOnceInstance() *singleSyncOnce {
    if singleOnceInstance != nil {
        once.Do(func() {
            fmt.Println("Creating single instance now.")
            singleOnceInstance = &singleSyncOnce{}
        })
    } else {
        fmt.Println("Single instance already created.")
    }
    return singleOnceInstance
}
