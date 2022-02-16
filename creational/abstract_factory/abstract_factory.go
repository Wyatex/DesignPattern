package abstract_factory

import "fmt"

type Lunch interface {
    Cook()
}

type Rise struct {
}

func (c *Rise) Cook() {
    fmt.Println("煮米饭")
}

type Vegetable struct {
}

func (c *Vegetable) Cook() {
    fmt.Println("煮青菜")
}

type LunchFactory interface {
    CreateFood() Lunch
    CreateVegetable() Lunch
}

type simpleLunchFactory struct {
}

func NewSimpleShapeFactory() LunchFactory {
    return &simpleLunchFactory{}
}

func (s *simpleLunchFactory) CreateFood() Lunch {
    return &Rise{}
}

func (s *simpleLunchFactory) CreateVegetable() Lunch {
    return &Vegetable{}
}
