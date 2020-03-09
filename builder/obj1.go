// Package builder ... Builder Design Pattern
package builder

// Obj1Builder ... How object1 should be built.
type Obj1Builder struct {
	obj *Object
}

// New ... Create a new object1.
func (o *Obj1Builder) New() BuildingProcess {
	o.obj = &Object{}
	return o
}

// SetNumber ... Set object1 number.
func (o *Obj1Builder) SetNumber() BuildingProcess {
	o.obj.number = 1
	return o
}

// SetName ... Set object1 name.
func (o *Obj1Builder) SetName() BuildingProcess {
	o.obj.name = "one"
	return o
}

// GetObject ... GetObject object1 itself.
func (o *Obj1Builder) GetObject() Object {
	return *o.obj
}
