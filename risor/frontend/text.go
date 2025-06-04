package frontend

import (
	"github.com/boxesandglue/boxesandglue/frontend"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const FrontendTextType = "frontend.text"

type text struct {
	Value *frontend.Text
}

// Type of the object.
func (txt *text) Type() object.Type {
	return FrontendTextType
}

// Inspect returns a string representation of the given object.
func (txt *text) Inspect() string {
	return txt.Value.String()
}

// Interface converts the given object to a native Go value.
func (txt *text) Interface() interface{} {
	return txt.Value
}

// Returns True if the given object is equal to this object.
func (txt *text) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (txt *text) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "settings":
		return &settings{txt: txt.Value}, true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (txt *text) SetAttr(name string, value object.Object) error {
	switch name {
	case "items":
		switch value.Type() {
		case object.LIST:
			lst := value.(*object.List).Value()
			for _, itm := range lst {
				switch t := itm.(type) {
				case *object.String:
					txt.Value.Items = append(txt.Value.Items, t.Value())
				case *text:
					txt.Value.Items = append(txt.Value.Items, t.Value)
				default:
					// fmt.Printf("~~> SetAttr/List/ itm %T\n", itm)
				}
			}
			return nil
		default:
			// fmt.Println("~~> not a list")
		}
	case "settings":
		// fmt.Println("~~> SetAttr/settings")
	}
	return object.Errorf("cannot set attribute %s on text", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (txt *text) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (txt *text) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on text", opType)
}

// Cost returns the incremental processing cost of this object.
func (txt *text) Cost() int {
	return 0
}
