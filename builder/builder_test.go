package builder

import "testing"

func TestBuilder(t *testing.T) {
	tests := []struct {
		number          int
		name            string
		buildingProcess BuildingProcess
	}{
		{1, "one", &Obj1Builder{}},
		{2, "two", &Obj2Builder{}},
	}

	director := Director{}

	for _, test := range tests {
		objBuilder := test.buildingProcess
		director.SetBuilder(objBuilder)
		director.Build()
		obj := objBuilder.GetObject()

		if obj.number != test.number {
			t.Errorf("Error. Want %v, got  %v.\n", test.number, obj.number)
		}

		if obj.name != test.name {
			t.Errorf("Error. Want %v, got %v.\n", test.name, obj.name)
		}
	}
}
