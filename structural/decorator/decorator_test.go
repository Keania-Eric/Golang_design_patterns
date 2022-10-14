package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling add ingredient of the pizza decorator it must return the text %s not '%s'", expectedText, pizzaResult)
	}
}

func TestOnionPizza_AddIngredient(t *testing.T) {
	onion := &OnionPizza{}
	onionResult, err := onion.AddIngredient()
	if err == nil {
		t.Errorf("when calling AddIngredient without an ingredientAdd on its field it should return an error and not %s", onionResult)
	}

	onion = &OnionPizza{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(onionResult, "onion") {
		t.Errorf("when calling AddIngredient with the onion decorator it must return a string with the word onion not %s", onionResult)
	}
	t.Log(onionResult)
}

func TestMeatPizza_AddIngredient(t *testing.T) {
	meat := &MeatPizza{}
	meatResult, err := meat.AddIngredient()
	if err == nil {
		t.Errorf("when calling AddIngredient without an ingredientAdd on its field it should return an error and not %s", meatResult)
	}

	meat = &MeatPizza{&PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(meatResult, "meat") {
		t.Errorf("when calling AddIngredient with the meat decorator it must return a string with the word meat not %s", meatResult)
	}
}

func TestPizzaDecorator_BothIngredients(t *testing.T) {
	pizza := &OnionPizza{&MeatPizza{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	expectedText := "Pizza with the following ingredients: meat onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When asking for a pizza with onion and meat the returned "+
			"string must contain the text '%s' but '%s' didn't have it",
			expectedText, pizzaResult)
	}
	t.Log(pizzaResult)
}
