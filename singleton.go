// Package singleton ... Singleton Design Pattern
package singleton

var uniqueSingleton *Singleton

// Singleton ... An example of a singleton struct. It should be unique.
type Singleton struct {
	value int
}

// NewSingleton ... Returns a singleton var. It should return the same
// singleton var every time it was called.
func NewSingleton() *Singleton {
	if uniqueSingleton == nil {
		uniqueSingleton = &Singleton{}
	}
	return uniqueSingleton
}

// Increment ... Increment the value of singleton to 2 at maximum.
func (s *Singleton) Increment() int {
	if s.value != 2 {
		s.value++
	}
	return s.value
}

// Decrement ... Decrement the value of singleton to 0 at minimum.
func (s *Singleton) Decrement() int {
	if s.value != 0 {
		s.value--
	}
	return s.value
}
