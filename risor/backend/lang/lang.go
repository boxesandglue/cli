package lang

import (
	"github.com/boxesandglue/boxesandglue/backend/lang"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type Lang struct {
	Value *lang.Lang
}

// Type of the object.
func (l *Lang) Type() object.Type {
	return "backend.lang"
}

// Inspect returns a string representation of the given object.
func (l *Lang) Inspect() string {
	return l.Value.Name
}

// Interface converts the given object to a native Go value.
func (l *Lang) Interface() interface{} {
	return l.Value
}

// Equals returns True if the given object is equal to this object.
func (l *Lang) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (l *Lang) GetAttr(name string) (object.Object, bool) {
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (l *Lang) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on lang", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (l *Lang) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (l *Lang) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("cannot run operation %s on lang", opType)
}

// Cost returns the incremental processing cost of this object.
func (l *Lang) Cost() int {
	return 0
}
