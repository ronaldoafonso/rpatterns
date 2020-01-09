package singleton

import "testing"

func TestGetSingleton(t *testing.T) {
	singleton1 := GetSingleton()
	if singleton1 == nil {
		t.Error("GetSingleton returning nil instead of a singleton.")
	}

	singleton2 := GetSingleton()
	if singleton1 != singleton2 {
		t.Error("GetSingleton returning not a unique singleton.")
	}

	valueOne := singleton1.Increment()
	if valueOne != 1 {
		t.Errorf("Increment should return 1 at first invocation, but returned %d.\n", valueOne)
	}

	valueTwo := singleton2.Increment()
	if valueTwo != 2 {
		t.Errorf("Increment should return 2 at second invocation, but returned %d.\n", valueTwo)
	}
}
