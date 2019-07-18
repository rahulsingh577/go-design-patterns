package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	counter int
}

func (s *singleton) AddOne() int {
	s.counter++
	return s.counter
}

var instance *singleton

func GetInstance() Singleton {

	if instance == nil {
		instance = new(singleton)
	}
	return instance
	//return singleton{counter: 0}
}
