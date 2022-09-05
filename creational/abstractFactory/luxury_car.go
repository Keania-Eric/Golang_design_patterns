package abstract_factory

type LuxuryCar struct{}

func (l *LuxuryCar) NumOfDoors() int {
	return 4
}

func (l *LuxuryCar) NumOfSeats() int {
	return 5
}

func (l *LuxuryCar) NumOfWheels() int {
	return 4
}
