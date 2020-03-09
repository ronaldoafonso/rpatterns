// Package builder ... Builder Design Pattern
package builder

// Object ... The object itself to be built.
type Object struct {
	number int
	name   string
}

// BuildingProcess ... The process of building an object.
type BuildingProcess interface {
	New() BuildingProcess
	SetNumber() BuildingProcess
	SetName() BuildingProcess
	GetObject() Object
}

// Director ... An example of a director struct.
type Director struct {
	builderProcess BuildingProcess
}

// SetBuilder ... Set the object building process.
func (d *Director) SetBuilder(b BuildingProcess) {
	d.builderProcess = b
	return
}

// Build ... Build the object using the process set before hand.
func (d *Director) Build() {
	d.builderProcess.New().SetNumber().SetName()
}
