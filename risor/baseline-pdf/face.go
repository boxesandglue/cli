package pdf

import (
	pdf "github.com/boxesandglue/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const TypeFace = "pdf.face"

type Face struct {
	Value *pdf.Face
}

// Type of the object.
func (face *Face) Type() object.Type {
	return TypeFace
}

// Inspect returns a string representation of the given object.
func (face *Face) Inspect() string {
	return face.Value.Filename
}

// Interface converts the given object to a native Go value.
func (face *Face) Interface() interface{} {
	return face.Value
}

// Equals returns True if the given object is equal to this object.
func (face *Face) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (face *Face) GetAttr(name string) (object.Object, bool) {
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (face *Face) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on text", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (face *Face) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (face *Face) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on Face", opType)
}

// Cost returns the incremental processing cost of this object.
func (face *Face) Cost() int {
	return 0
}
