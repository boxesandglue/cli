package pdf

import (
	"context"
	"fmt"
	"io"
	"os"

	pdf "github.com/boxesandglue/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

// TypePDF is the type of the PDF object.
const TypePDF = "pdf.pdf"

// PDF is the object that represents a PDF document.
// It is a wrapper around the pdf.PDF type from the baseline-pdf package.
type PDF struct {
	Value *pdf.PDF
}

func newBaselinePDF(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("pdf.new() takes exactly one argument (filename)")
	}
	var w io.Writer
	firstArg := args[0]
	switch firstArg.Type() {
	case object.STRING:
		filename := firstArg.(*object.String).Value()
		if filename == "" {
			return object.ArgsErrorf("pdf.new() expects a non-empty string argument (filename)")
		}
		// Load the PDF file using the baseline-pdf package.
		f, err := os.Create(filename) // Ensure the file exists, this is just a placeholder.
		if err != nil {
			return object.Errorf("failed to create PDF file: %v", err)
		}
		defer f.Close()
		w = f
	case object.FILE:
		w = firstArg.(*object.File).Value()
	default:
		fmt.Println(`~~> firstArg`, firstArg.Type())
	}

	return &PDF{
		Value: pdf.NewPDFWriter(w),
	}
}

// pdfNewFace creates a new Face object from the given filename and index. The
// index is optional. If it is not provided, the default index of 0 is used.
func (pdf *PDF) pdfNewFace(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return object.ArgsErrorf("pdf.new_face() takes one or two arguments (filename, index)")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("pdf.new_face() expects a string argument (filename)")
	}
	idx := 0
	if len(args) == 2 {
		if args[1].Type() == object.INT {
			idx = int(args[1].(*object.Int).Value())
		}
	}

	filename := args[0].(*object.String).Value()
	f, err := pdf.Value.LoadFace(filename, idx)
	if err != nil {
		return object.NewError(err)
	}
	return &Face{Value: f}
}

// Type of the object.
func (pdf *PDF) Type() object.Type {
	return TypePDF
}

// Inspect returns a string representation of the given object.
func (pdf *PDF) Inspect() string {
	return "PDF file"
}

// Interface converts the given object to a native Go value.
func (pdf *PDF) Interface() interface{} {
	return pdf.Value
}

// Equals returns True if the given object is equal to this object.
func (pdf *PDF) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (pdf *PDF) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "new_face":
		return object.NewBuiltin("pdf.new_face", pdf.pdfNewFace), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (pdf *PDF) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on pdf", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (pdf *PDF) IsTruthy() bool {
	return false
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (pdf *PDF) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on pdf", opType)
}

// Cost returns the incremental processing cost of this object.
func (pdf *PDF) Cost() int {
	return 0
}

// Module returns the frontend module.
func Module() *object.Module {
	return object.NewBuiltinsModule("baselinepdf", map[string]object.Object{
		"new": object.NewBuiltin("new", newBaselinePDF),
	})
}
