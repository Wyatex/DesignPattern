package factory

import "testing"

func TestNewRestaurant(t *testing.T) {
    NewRestaurant("mcl").GetFood()
    NewRestaurant("mcd").GetFood()
}
