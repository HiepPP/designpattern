package prototype

import (
	"errors"
	"fmt"
)

// Mục đích của Prototype pattern là have an object or set object that is already created at compilation time.
// Nhưng có thể clone nhiều lần ở runtime.
// Giúp avoid repetitive object creation

// Maintain a set of objects that will be cloned to create new instances
// Provide a default value of some type to start working on top of it
// Free CPU of complex object initialization to take more memory resources

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtCache struct{}

func (s *ShirtCache) GetClone(num int) (ItemInfoGetter, error) {
	switch num {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}
type ShirtColor byte
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Colro id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func GetShirtCloner() ShirtCloner {
	return new(ShirtCache)
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

func (s *Shirt) GetPrice() float32 {
	return s.Price
}
