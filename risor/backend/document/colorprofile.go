package document

import (
	"github.com/boxesandglue/boxesandglue/backend/document"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const TypeColorProfile = "colorprofile"

type ColorProfile struct {
	Value *document.ColorProfile
}

// Type of the object.
func (cp *ColorProfile) Type() object.Type {
	return TypeColorProfile
}

// Inspect returns a string representation of the given object.
func (cp *ColorProfile) Inspect() string {
	return cp.Value.Identifier
}

// Interface converts the given object to a native Go value.
func (cp *ColorProfile) Interface() interface{} {
	return cp.Value
}

// Returns True if the given object is equal to this object.
func (cp *ColorProfile) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (cp *ColorProfile) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "identifier":
		return object.NewString(cp.Value.Identifier), true
	case "registry":
		return object.NewString(cp.Value.Registry), true
	case "info":
		return object.NewString(cp.Value.Info), true
	case "condition":
		return object.NewString(cp.Value.Condition), true
	case "colors":
		return object.NewInt(int64(cp.Value.Colors)), true
	default:
		return nil, false
	}
}

// SetAttr sets the attribute with the given name on this object.
func (cp *ColorProfile) SetAttr(name string, value object.Object) error {
	switch name {
	case "identifier":
		if v, ok := value.(*object.String); ok {
			cp.Value.Identifier = v.Value()
			return nil
		}
	case "registry":
		if v, ok := value.(*object.String); ok {
			cp.Value.Registry = v.Value()
			return nil
		}
	case "info":
		if v, ok := value.(*object.String); ok {
			cp.Value.Info = v.Value()
			return nil
		}
	case "condition":
		if v, ok := value.(*object.String); ok {
			cp.Value.Condition = v.Value()
			return nil
		}
	case "colors":
		if v, ok := value.(*object.Int); ok {
			cp.Value.Colors = int(v.Value())
			return nil
		}
	default:
		return object.Errorf("cannot set attribute %s on color profile", name)
	}
	return object.Errorf("invalid type for attribute %s", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (cp *ColorProfile) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (cp *ColorProfile) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on color profile", opType)
}

// Cost returns the incremental processing cost of this object.
func (cp *ColorProfile) Cost() int {
	return 0
}
