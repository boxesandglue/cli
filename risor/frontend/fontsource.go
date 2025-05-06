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
	return fs
}

// Returns True if the given object is equal to this object.
func (fs *fontSource) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (fs *fontSource) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "location":
		return object.NewString(fs.location), true
	case "name":
		return object.NewString(fs.name), true
	case "fontFeatures":
		l := object.NewList(nil)
		for _, feature := range fs.fontFeatures {
			l.Append(object.NewString(feature))
		}
		return l, true
	case "index":
		return object.NewInt(int64(fs.index)), true
	default:
		return nil, false
	}
}

// SetAttr sets the attribute with the given name on this object.
func (fs *fontSource) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on fontsource", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (fs *fontSource) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (fs *fontSource) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on fontsource", opType)
}

// Cost returns the incremental processing cost of this object.
func (fs *fontSource) Cost() int {
	return 0
}
