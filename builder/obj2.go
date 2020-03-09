// Package builder ... Builder Design Pattern
package builder

// Obj2Builder ... How object2 should be built.
type Obj2Builder struct {
	obj *Object
}

// New ... Create a new object2.
func (o *Obj2Builder) New() BuildingProcess {
	o.obj = &Object{}
	return o
}

// SetNumber ... Set object2 number.
func (o *Obj2Builder) SetNumber() BuildingProcess {
	o.obj.number = 2
	return o
}

// SetName ... Set object2 name.
func (o *Obj2Builder) SetName() BuildingProcess {
	o.obj.name = "two"
	return o
}

// GetObject ... GetObject object2 itself.
func (o *Obj2Builder) GetObject() Object {
	return *o.obj
}
