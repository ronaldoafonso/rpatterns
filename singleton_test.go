package singleton

import "testing"

func TestSingleton(t *testing.T) {
	s1 := NewSingleton()
	if s1 == nil {
		t.Error("NewSingleton returning nil instead of a singleton.")
	}

	for i := 0; i < 10; i++ {
		if s := NewSingleton(); s != s1 {
			t.Error("NewSingleton returning not a unique singleton.")
		}
	}

	s2 := NewSingleton()
	tests := []struct {
		name string
		do   func() int
		want int
	}{
		{"s1.Increment", s1.Increment, 1},
		{"s2.Increment", s2.Increment, 2},
		{"s1.Decrement", s1.Decrement, 1},
		{"s2.Decrement", s2.Decrement, 0},
		{"s1.Decrement", s1.Decrement, 0},
		{"s2.Increment", s2.Increment, 1},
		{"s1.Increment", s1.Increment, 2},
		{"s2.Increment", s2.Increment, 2},
	}

	for n, test := range tests {
		got := test.do()
		if got != test.want {
			t.Errorf("Test %v - Function: %v error. Want: %v, got: %v.",
				n+1,
				test.name,
				test.want,
				got)
		}
	}
}
