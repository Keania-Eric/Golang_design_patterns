package abstract_factory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	MotorBikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := MotorBikeF.Build(SportMotorBikeType)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motobike Vehicles has %d number of wheels", motorbikeVehicle.NumOfWheels())

	sportBike, ok := motorbikeVehicle.(MotorBike)

	if !ok {
		t.Fatal("Struct assertion failed")
	}

	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())
}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)

	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car Vehicle has %d wheels", carVehicle.NumOfWheels())

	luxuryCar, ok := carVehicle.(Car)

	if !ok {
		t.Fatal("Struct assertion failed")
	}

	t.Logf("luxury cars has %d wheels \n", luxuryCar.NumOfSeats())
}
