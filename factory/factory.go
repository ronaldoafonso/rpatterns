// Package factory ... Factory Design Pattern
package factory

import (
	"errors"
)

// ObjType ... Type for objects.
type ObjType int

const (
	nonExistentType ObjType = iota
	type1
	type2
)

// ObjMethodInterface ... Which interface must be implemented by objects.
type ObjMethodInterface interface {
	Method(string) string
}

// GetObjectMethod ... Main method for the factory pattern.
func GetObjectMethod(objType ObjType) (ObjMethodInterface, error) {
	var obj ObjMethodInterface
	var err error

	switch objType {
	case type1:
		obj = Obj1{}
	case type2:
		obj = Obj2{}
	default:
		err = errors.New("Non Existent Type")
	}

	return obj, err
}
