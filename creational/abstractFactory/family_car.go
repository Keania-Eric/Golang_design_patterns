package abstract_factory

type FamilyCar struct{}

func (f *FamilyCar) NumOfDoors() int {
	return 5
}

func (f *FamilyCar) NumOfSeats() int {
	return 5
}

func (l *FamilyCar) NumOfWheels() int {
	return 4
}
