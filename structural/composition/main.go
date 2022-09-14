package main

import "fmt"

type Swimmer interface {
	Swim()
}

type SwimmerImpl struct{}

func (swl *SwimmerImpl) Swim() {
	fmt.Println("Swimming!")
}

type Trainer interface {
	Train()
}
type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type Animal struct{}

func (an *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal // watch as we embed this year a way of doing inheritance in go
	Swim   func()
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

type CompositeSwimmerB struct {
	Swimmer
	Trainer
}

func Swim() {
	fmt.Println("Swimming!")
}

func main() {
	fish := Shark{Swim: Swim}
	fish.Eat()
	fish.Swim()
	localSwim := Swim

	swimmer := CompositeSwimmerA{MySwim: localSwim}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	// swimmer2 := CompositeSwimmerB{
	// 	&Athlete{},
	// 	&SwimmerImpl{},
	// }

	// swimmer2.Train()
	// swimmer2.Swim()
}
