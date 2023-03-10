package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}

	return car
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	track := Track{distance: distance}

	return track
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	track := NewTrack(car.distance)

	if car.battery-car.batteryDrain >= 0 {
		car.battery = car.battery - car.batteryDrain
		track.distance += car.speed
	}

	car.distance = track.distance

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	drivenTimes := car.battery / car.batteryDrain
	maxDist := car.speed * drivenTimes

	if maxDist >= track.distance {
		return true
	} else {
		return false
	}

}
