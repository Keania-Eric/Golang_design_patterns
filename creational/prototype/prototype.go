package prototype

import (
	"errors"
	"fmt"
)

type ShirtRepository interface {
	GetClone(s int) (ItemInfoGetter, error)
}

type ItemInfoGetter interface {
	GetInfo() string
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

func GetShirtCloner() ShirtRepository {
	return new(ShirtStore)
}

type ShirtStore struct{} // implements shirts repository
func (sh *ShirtStore) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		item := *whitePrototype
		return &item, nil
	case Black:
		item := *blackPrototype
		return &item, nil
	case Blue:
		item := *bluePrototype
		return &item, nil
	default:
		return nil, errors.New("shirt model not recognized")
	}

}

type ShirtColor byte

type Shirt struct { // implements ItemInfoGetter
	Price float32
	SKU   string
	Color ShirtColor
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU %s and color id %d costs %f", s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}
