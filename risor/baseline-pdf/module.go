package pdf

import (
	"context"
	"io"
	"os"

	rpdf "github.com/boxesandglue/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

// TypePDF is the type of the PDF object.
const TypePDF = "pdf.pdf"

// PDF is the object that represents a PDF document.
// It is a wrapper around the pdf.PDF type from the baseline-pdf package.
type PDF struct {
	Value *rpdf.PDF
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
		w = f
	case object.FILE:
		w = firstArg.(*object.File).Value()
	default:
		// fmt.Println(`~~> firstArg`, firstArg.Type())
	}

	return &PDF{
		Value: rpdf.NewPDFWriter(w),
	}
}

// addPage adds a new page to the PDF document. It takes two arguments: the
// stream object and the page number. If the page number is not provided, it
// defaults to the next available object number.
func (pdf *PDF) addPage(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return object.ArgsErrorf("pdf.add_page() takes one or two arguments (stream, [object number])")
	}

	if args[0].Type() != TypeObject {
		return object.ArgsErrorf("pdf.add_page() expects a pdf.object as first argument (stream)")
	}
	streamObj := args[0].(*Object)
	var pageObject rpdf.Objectnumber
	if len(args) == 2 {
		if args[1].Type() != object.INT {
			return object.ArgsErrorf("pdf.add_page() expects an int as second argument (object number)")
		}
	} else {
		pageObject = rpdf.Objectnumber(0)
	}
	pdfPage := pdf.Value.AddPage(streamObj.Value, pageObject)
	return &Page{Value: pdfPage}
}

// finish finalizes the PDF document and writes it to the underlying writer.
func (pdf *PDF) finish(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.ArgsErrorf("pdf.finish() takes no arguments")
	}
	if err := pdf.Value.FinishAndClose(); err != nil {
		return object.NewError(err)
	}
	return object.Nil
}

// pdfPrint writes the given string to the PDF document.
func (pdf *PDF) pdfPrint(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("pdf.print() takes exactly one argument (string)")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("pdf.print() expects a string argument")
	}
	str := args[0].(*object.String).Value()
	if err := pdf.Value.Print(str); err != nil {
		return object.NewError(err)
	}
	return object.Nil
}

// pdfPrintf writes the formatted string to the PDF document.
func (pdf *PDF) pdfPrintf(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 {
		return object.ArgsErrorf("pdf.printf() takes at least one argument (format string)")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("pdf.printf() expects a string as first argument (format string)")
	}
	format := args[0].(*object.String).Value()
	var formatArgs []any
	for _, arg := range args[1:] {
		formatArgs = append(formatArgs, arg.Interface())
	}
	if err := pdf.Value.Printf(format, formatArgs...); err != nil {
		return object.NewError(err)
	}
	return object.Nil
}

// pdfPrintln writes the given string to the PDF document, followed by a newline.
func (pdf *PDF) pdfPrintln(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("pdf.println() takes exactly one argument (string)")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("pdf.println() expects a string argument")
	}
	str := args[0].(*object.String).Value()
	if err := pdf.Value.Println(str); err != nil {
		return object.NewError(err)
	}
	return object.Nil
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

func (pdf *PDF) pdfNewObject(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.ArgsErrorf("pdf.new_object() takes no arguments")
	}
	object := pdf.Value.NewObject()
	return &Object{Value: object}
}

func (pdf *PDF) pdfLoadImageFile(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 || len(args) > 3 {
		return object.ArgsErrorf("pdf.load_image_file() takes one to three arguments (filename, [box], [pagenumber])")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("pdf.load_image_file() expects a string as first argument (filename)")
	}
	filename := args[0].(*object.String).Value()
	box := ""
	pagenumber := 1
	if len(args) >= 2 {
		if args[1].Type() != object.STRING {
			return object.ArgsErrorf("pdf.load_image_file() expects a string as second argument (box)")
		}
		box = args[1].(*object.String).Value()
	}
	if len(args) == 3 {
		if args[2].Type() != object.INT {
			return object.ArgsErrorf("pdf.load_image_file() expects an int as third argument (pagenumber)")
		}
		pagenumber = int(args[2].(*object.Int).Value())
	}
	imgfile, err := pdf.Value.LoadImageFileWithBox(filename, box, pagenumber)
	if err != nil {
		return object.NewError(err)
	}
	return &ImageFile{Value: imgfile}
}

func (pdf *PDF) pdfNewObjectWithNumber(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("pdf.new_object_with_number() takes exactly one argument (object number)")
	}
	switch args[0].Type() {
	case object.INT:
		objNum := rpdf.Objectnumber(args[0].(*object.Int).Value())
		object := pdf.Value.NewObjectWithNumber(objNum)
		return &Object{Value: object}
	case typeObjectNumber:
		objNum := args[0].(*objectNumber).Value
		object := pdf.Value.NewObjectWithNumber(objNum)
		return &Object{Value: object}
	default:
		return object.ArgsErrorf("pdf.new_object_with_number() expects an int argument (object number)")
	}
}

func (pdf *PDF) pdfNextObject(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.ArgsErrorf("pdf.next_object() takes no arguments")
	}
	nextObjNum := pdf.Value.NextObject()
	return objectNumber{Value: nextObjNum}
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
	case "add_page":
		return object.NewBuiltin("pdf.add_page", pdf.addPage), true
	case "default_page_height":
		return object.NewFloat(pdf.Value.DefaultPageHeight), true
	case "default_page_width":
		return object.NewFloat(pdf.Value.DefaultPageWidth), true
	case "default_offset_x":
		return object.NewFloat(pdf.Value.DefaultOffsetX), true
	case "default_offset_y":
		return object.NewFloat(pdf.Value.DefaultOffsetY), true
	case "finish":
		return object.NewBuiltin("pdf.finish", pdf.finish), true
	case "load_image_file":
		return object.NewBuiltin("pdf.load_image_file", pdf.pdfLoadImageFile), true
	case "new_face":
		return object.NewBuiltin("pdf.new_face", pdf.pdfNewFace), true
	case "new_object":
		return object.NewBuiltin("pdf.new_object", pdf.pdfNewObject), true
	case "new_object_with_number":
		return object.NewBuiltin("pdf.new_object_with_number", pdf.pdfNewObjectWithNumber), true
	case "next_object":
		return object.NewBuiltin("pdf.next_object", pdf.pdfNextObject), true
	case "print":
		return object.NewBuiltin("pdf.print", pdf.pdfPrint), true
	case "printf":
		return object.NewBuiltin("pdf.printf", pdf.pdfPrintf), true
	case "println":
		return object.NewBuiltin("pdf.println", pdf.pdfPrintln), true
	case "size":
		return object.NewInt(pdf.Value.Size()), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (pdf *PDF) SetAttr(name string, value object.Object) error {
	switch name {
	case "default_page_height":
		switch value.Type() {
		case object.INT:
			pdf.Value.DefaultPageHeight = float64(value.(*object.Int).Value())
			return nil
		case object.FLOAT:
			pdf.Value.DefaultPageHeight = value.(*object.Float).Value()
			return nil
		default:
			return object.Errorf("expected int or float for default_page_height, got %s", value.Type())
		}
	case "default_page_width":
		switch value.Type() {
		case object.INT:
			pdf.Value.DefaultPageWidth = float64(value.(*object.Int).Value())
			return nil
		case object.FLOAT:
			pdf.Value.DefaultPageWidth = value.(*object.Float).Value()
			return nil
		default:
			return object.Errorf("expected int or float for default_page_width, got %s", value.Type())
		}
	case "default_offset_x":
		switch value.Type() {
		case object.INT:
			pdf.Value.DefaultOffsetX = float64(value.(*object.Int).Value())
			return nil
		case object.FLOAT:
			pdf.Value.DefaultOffsetX = value.(*object.Float).Value()
			return nil
		default:
			return object.Errorf("expected int or float for default_offset_x, got %s", value.Type())
		}
	case "default_offset_y":
		switch value.Type() {
		case object.INT:
			pdf.Value.DefaultOffsetY = float64(value.(*object.Int).Value())
			return nil
		case object.FLOAT:
			pdf.Value.DefaultOffsetY = value.(*object.Float).Value()
			return nil
		default:
			return object.Errorf("expected int or float for default_offset_y, got %s", value.Type())
		}
	}
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
