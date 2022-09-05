package abstract_factory

type CruiseMotorBike struct{}

func (cb *CruiseMotorBike) NumOfWheels() int {
	return 2
}

func (cb *CruiseMotorBike) NumOfSeats() int {
	return 2
}

func (cb *CruiseMotorBike) GetMotorbikeType() int {
	return CruiseMotorBikeType
}
