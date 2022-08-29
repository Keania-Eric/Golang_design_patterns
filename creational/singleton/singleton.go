package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

// In Go, you can initialize a pointer to a struct as nil,
// but you cannot initialize a struct to nil (the equivalent of NULL).
//  So the var instance *singleton line defines a pointer to a struct of type Singleton as nil, and the variable called instance.

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
