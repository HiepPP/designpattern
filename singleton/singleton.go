package singleton

// We need a single, shared value, of some particular type
// We need restrict object creation of some type to a single unit along the entire program

import "sync"

var (
	once sync.Once
)

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

//-------------------------

type OnlyYouInterface interface {
	AddOther() int
}

type OnlyYou struct {
	MeInYou int
}

var onlyInstance *OnlyYou

func GetOnlyYou() OnlyYouInterface {
	once.Do(func() {
		onlyInstance = new(OnlyYou)
	})
	return onlyInstance
}

func (o *OnlyYou) AddOther() int{
	o.MeInYou++
	return o.MeInYou
}
