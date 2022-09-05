package abstract_factory

import (
	"fmt"
)

type CarFactory struct{}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("car with id %d not recognized", v)
	}
}
