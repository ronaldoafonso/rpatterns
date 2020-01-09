package singleton

type Singleton struct {
	value int
}

var uniqueSingleton *Singleton

func GetSingleton() *Singleton {
	if uniqueSingleton == nil {
		uniqueSingleton = new(Singleton)
	}
	return uniqueSingleton
}

func (s *Singleton) Increment() int {
	s.value++
	return s.value
}
