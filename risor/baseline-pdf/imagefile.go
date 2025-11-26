package pdf

import (
	"context"

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
	case "close":
		return object.NewBuiltin("pdf.imagefile.close", imgf.close), true
	case "get_pdf_box_dimensions":
		return object.NewBuiltin("pdf.imagefile.get_pdf_box_dimensions", imgf.getPDFBoxDimensions), true
	case "internal_name":
		return object.NewString(imgf.Value.InternalName()), true
	case "page_number":
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

// implementation

func (imgf *ImageFile) close(ctx context.Context, args ...object.Object) object.Object {
	err := imgf.Value.Close()
	if err != nil {
		return object.Errorf("error closing image file: %s", err)
	}
	return object.Nil
}

func (imgf *ImageFile) getPDFBoxDimensions(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.Errorf("expected 2 arguments, got %d", len(args))
	}
	pagenumberObj := args[0]
	boxNameObj := args[1]

	pagenumberInt, ok := pagenumberObj.(*object.Int)
	if !ok {
		return object.TypeErrorf("expected int for page number, got %s", pagenumberObj.Type())
	}
	boxNameStr, ok := boxNameObj.(*object.String)
	if !ok {
		return object.TypeErrorf("expected string for box name, got %s", boxNameObj.Type())
	}

	pagenumber := int(pagenumberInt.Value())
	boxName := boxNameStr.String()

	dimensions, err := imgf.Value.GetPDFBoxDimensions(pagenumber, boxName)
	if err != nil {
		return object.Errorf("error getting PDF box dimensions: %s", err)
	}

	result := object.NewMap(nil)
	for k, v := range dimensions {
		result.Set(k, object.NewFloat(v))
	}
	return result
}
