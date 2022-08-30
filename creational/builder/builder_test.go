package builder

import "testing"

func TestBuilderPattern(t *testing.T) {

	manufacturer := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturer.SetBuilder(carBuilder)
	manufacturer.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure of a car must be 'Car' and they were %s", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturer.SetBuilder(bikeBuilder)
	manufacturer.Construct()

	motorBike := bikeBuilder.GetVehicle()

	if motorBike.Wheels != 2 {
		t.Errorf("Wheels on a motor bike must be 2 and they were %d", motorBike.Wheels)
	}

	if motorBike.Structure != "Bike" {
		t.Errorf("Structure of a motor bike must be 'Bike' and they were %s", motorBike.Structure)
	}

	if motorBike.Seats != 1 {
		t.Errorf("Seats on a motot bike must be 1 and they were %d", motorBike.Seats)
	}

}
