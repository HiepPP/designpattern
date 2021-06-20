package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
	expectedCounter := counter1
	currentCounter := counter1.AddOne()
	if currentCounter != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCounter)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	currentCounter = counter2.AddOne()
	if currentCounter != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCounter)
	}
}

func TestOnlyYou(t *testing.T) {
	onlyYou := GetOnlyYou()
	if onlyYou == nil {
		t.Error("Onlyyou error")
	}

	result1 := onlyYou.AddOther()
	if result1 != 1 {
		t.Errorf("Sai roi, phai la 1")
	}

	result2 := onlyYou.AddOther()
	if result2 != 2{
		t.Error("Result must be 2")
	}
}
