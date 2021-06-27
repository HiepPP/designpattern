package prototype

import "testing"

func TestGetShirtCloner(t *testing.T) {
	shirtCache := GetShirtCloner()
	//First case
	if shirtCache == nil{
		t.Fatal("Received cache was nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil{
		t.Error(err)
	}
    // second case
	if item1 == whitePrototype{
		t.Error("item1 cannot be equal to the white prototype")
	}
	// Third case
	shirt1, ok := item1.(*Shirt)
	if !ok{
		t.Fatal("Type assertion for shirt1 could not be done successfully")
	}
	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil{
		t.Fatal(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok{
		t.Fatal("Type assertion for shirt2 could not be done successfully")
	}

	if shirt1.SKU == shirt2.SKU{
		t.Error("SKU of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2{
		t.Error("Shirt1 cannot be equal to Shirt2")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())
	t.Logf("LOG: the memory position of the shirts are different %p != %p \n\n", &shirt1, &shirt2)
}