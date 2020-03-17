package factory

// Obj1 ... Object1 implementation.
type Obj1 struct{}

// Method ... Base method for factory pattern.
func (o Obj1) Method(msg string) string {
	return msg + " returned"
}
