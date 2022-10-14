package decorator

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type MeatPizza struct {
	Ingredient IngredientAdd
}

func (m *MeatPizza) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("an Ingredient is needed in the ingredient field of meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", s, "meat"), nil
}

type OnionPizza struct {
	Ingredient IngredientAdd
}

func (o *OnionPizza) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("an Ingredient is needed in the ingredient field of the onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", s, "onion"), nil
}
