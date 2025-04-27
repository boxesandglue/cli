package frontend

import (
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type fontSource struct {
	location     string
	name         string
	fontFeatures []string
	index        int
}

// Type of the object.
func (fs *fontSource) Type() object.Type {
	return "frontend.fontsource"
}

// Inspect returns a string representation of the given object.
func (fs *fontSource) Inspect() string {
	return fs.location
}

// Interface converts the given object to a native Go value.
func (fs *fontSource) Interface() interface{} {
	panic("not implemented") // TODO: Implement
}

// Returns True if the given object is equal to this object.
func (fs *fontSource) Equals(other object.Object) object.Object {
	panic("not implemented") // TODO: Implement
}

// GetAttr returns the attribute with the given name from this object.
func (fs *fontSource) GetAttr(name string) (object.Object, bool) {
	panic("not implemented") // TODO: Implement
}

// SetAttr sets the attribute with the given name on this object.
func (fs *fontSource) SetAttr(name string, value object.Object) error {
	panic("not implemented") // TODO: Implement
}

// IsTruthy returns true if the object is considered "truthy".
func (fs *fontSource) IsTruthy() bool {
	panic("not implemented") // TODO: Implement
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (fs *fontSource) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	panic("not implemented") // TODO: Implement
}

// Cost returns the incremental processing cost of this object.
func (fs *fontSource) Cost() int {
	panic("not implemented") // TODO: Implement
}
