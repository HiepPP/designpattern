package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirection{}
	//--------------------------car----------
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()
	if car.Wheels != 4 {
		t.Errorf("Wheels on car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}
	//----------------------bike-------------

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motoBike := bikeBuilder.GetVehicle()
	motoBike.Seats = 1
	if motoBike.Wheels != 2 {
		t.Errorf("Wheels on motobike must be 2 and they were %d\n", motoBike.Wheels)
	}
	if motoBike.Structure != "MotoBike" {
		t.Errorf("Structure on a motobike must be 'MotoBike' and was %s\n", motoBike.Structure)
	}
}
