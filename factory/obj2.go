package factory

// Obj2 ... Objject2 implementation.
type Obj2 struct{}

// Method ... Base method for factory pattern.
func (o Obj2) Method(msg string) string {
	return msg + " returned"
}
