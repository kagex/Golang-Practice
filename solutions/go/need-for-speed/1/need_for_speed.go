package speed

type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

type Track struct {
    distance int
}
// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {
        speed: speed,
        battery: 100,
        batteryDrain: batteryDrain,
    }
}


// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track {
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
        car.distance += car.speed
        car.battery -= car.batteryDrain
    }
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	drivenCount := car.battery / car.batteryDrain
    if drivenCount * car.speed >= track.distance {
        return true
    }
    return false
}
