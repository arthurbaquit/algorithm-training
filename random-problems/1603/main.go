package main

type ParkingSystem struct {
    MaxSpots map[int]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    var ps ParkingSystem
    ps.MaxSpots = make(map[int]int)
    ps.MaxSpots[1] = big
    ps.MaxSpots[2] = medium
    ps.MaxSpots[3] = small
    return ps
}


func (ps *ParkingSystem) AddCar(carType int) bool {
    if ps.MaxSpots[carType] == 0 {return false}
    ps.MaxSpots[carType]--
    return true
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */