// Package factory 工厂模式
package factory

import "fmt"

type Restaurant interface {
    GetFood()
}

type Michelin struct {
}

func (m *Michelin) GetFood() {
    fmt.Println("米其林三星饭菜制作完成")
}

type McDonald struct {
}

func (m *McDonald) GetFood() {
    fmt.Println("麦当劳快餐制作完成")
}

func NewRestaurant(name string) Restaurant {
    switch name {
    case "mcd":
        return &McDonald{}
    case "mcl":
        return &Michelin{}
    }
    return nil
}
