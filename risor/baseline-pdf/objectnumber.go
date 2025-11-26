package pdf

import (
	rpdf "github.com/boxesandglue/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

// typeObjectNumber is the type name for PDF object number.
const typeObjectNumber = "baseline-pdf.objectnumber"

// objectNumber represents a PDF object number in Risor.
type objectNumber struct {
	Value rpdf.Objectnumber
}

// Type of the object.
func (onum objectNumber) Type() object.Type {
	return typeObjectNumber
}

// Inspect returns a string representation of the given object.
func (onum objectNumber) Inspect() string {
	return onum.Value.String()
}

// Interface converts the given object to a native Go value.
func (onum objectNumber) Interface() any {
	return onum.Value
}

// Equals returns True if the given object is equal to this object.
func (onum objectNumber) Equals(other object.Object) object.Object {
	return object.NewBool(onum.Value == other.(*objectNumber).Value)
}

// GetAttr returns the attribute with the given name from this object.
func (onum objectNumber) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "ref":
		return object.NewString(onum.Value.Ref()), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (onum objectNumber) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on text", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (onum objectNumber) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (onum objectNumber) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on Object", opType)
}

// Cost returns the incremental processing cost of this object.
func (onum objectNumber) Cost() int {
	return 0
}
