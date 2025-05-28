package pdf

import (
	pdf "github.com/boxesandglue/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const PDFDocumentType = "pdf.document"

type PDFDocument struct {
	// The filename of the PDF file
	Value *pdf.PDF
}

// Type of the object.
func (pdf *PDFDocument) Type() object.Type {
	return PDFDocumentType
}

// Inspect returns a string representation of the given object.
func (pdf *PDFDocument) Inspect() string {
	return PDFDocumentType
}

// Interface converts the given object to a native Go value.
func (pdf *PDFDocument) Interface() interface{} {
	return pdf.Value
}

// Equals returns True if the given object is equal to this object.
func (pdf *PDFDocument) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (pdf *PDFDocument) GetAttr(name string) (object.Object, bool) {
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (pdf *PDFDocument) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on pdf.document", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (pdf *PDFDocument) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (pdf *PDFDocument) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on pdf.document", opType)
}

// Cost returns the incremental processing cost of this object.
func (pdf *PDFDocument) Cost() int {
	return 0
}
