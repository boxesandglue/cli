package pdf

import (
	pdf "github.com/boxesandglue/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

// ImageFile represents a PDF image file object.
type ImageFile struct {
	Value *pdf.Imagefile
}

// Type of the object.
func (imgf *ImageFile) Type() object.Type {
	return "baseline-pdf.imagefile"
}

// Inspect returns a string representation of the given object.
func (imgf *ImageFile) Inspect() string {
	return imgf.Value.Filename
}

// Interface converts the given object to a native Go value.
func (imgf *ImageFile) Interface() interface{} {
	return imgf.Value
}

// Equals returns True if the given object is equal to this object.
func (imgf *ImageFile) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (imgf *ImageFile) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "pagenumber":
		return object.NewInt(int64(imgf.Value.PageNumber)), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (imgf *ImageFile) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on font", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (imgf *ImageFile) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (imgf *ImageFile) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on ImageFile", opType)
}

// Cost returns the incremental processing cost of this object.
func (imgf *ImageFile) Cost() int {
	return 0
}
