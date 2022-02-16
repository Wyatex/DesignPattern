package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
    factory := NewSimpleShapeFactory()
    food := factory.CreateFood()
    food.Cook()

    vegetable := factory.CreateVegetable()
    vegetable.Cook()
}
