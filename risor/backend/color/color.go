package color

import (
	"github.com/boxesandglue/boxesandglue/backend/color"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const BackendColorType = "backend.color"

type RColor struct {
	Value *color.Color
}

// Type of the object.
func (col *RColor) Type() object.Type {
	return BackendColorType
}

// Inspect returns a string representation of the given object.
func (col *RColor) Inspect() string {
	return col.Value.String()
}

// Interface converts the given object to a native Go value.
func (col *RColor) Interface() any {
	return col.Value
}

// Equals returns True if the given object is equal to this object.
func (col *RColor) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (col *RColor) GetAttr(name string) (object.Object, bool) {
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (col *RColor) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on backend.Color", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (col *RColor) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (col *RColor) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on backend.Color", opType)
}

// Cost returns the incremental processing cost of this object.
func (col *RColor) Cost() int {
	return 0
}
