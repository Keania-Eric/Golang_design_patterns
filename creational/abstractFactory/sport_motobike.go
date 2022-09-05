package abstract_factory

type SportMotorBike struct{}

func (sb *SportMotorBike) NumOfWheels() int {
	return 2
}

func (sb *SportMotorBike) NumOfSeats() int {
	return 1
}

func (sb *SportMotorBike) GetMotorbikeType() int {
	return SportMotorBikeType
}
