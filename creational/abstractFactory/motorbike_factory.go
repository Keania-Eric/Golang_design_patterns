package abstract_factory

import (
	"fmt"
)

type MotorbikeFactory struct{}

const (
	SportMotorBikeType  = 1
	CruiseMotorBikeType = 2
)

func (mb *MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorBikeType:
		return new(SportMotorBike), nil
	case CruiseMotorBikeType:
		return new(CruiseMotorBike), nil
	default:
		return nil, fmt.Errorf("motor bike with id %d not recognized", v)
	}
}
