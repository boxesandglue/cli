package document

import (
	"context"

	"github.com/boxesandglue/boxesandglue/backend/document"
	"github.com/boxesandglue/boxesandglue/frontend"
	rlang "github.com/boxesandglue/cli/risor/backend/lang"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type Document struct {
	PDFDoc *document.PDFDocument
}

func (doc *Document) newPage(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.ArgsErrorf("document.new_page() takes no arguments")
	}
	return newPage(doc.PDFDoc)
}

func (doc *Document) finish(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.ArgsErrorf("document.finish() takes no arguments")
	}
	doc.PDFDoc.Finish()
	return nil
}

// Type of the object.
func (doc *Document) Type() object.Type {
	return "backend.document"
}

// Inspect returns a string representation of the given object.
func (doc *Document) Inspect() string {
	return doc.PDFDoc.Filename
}

// Interface converts the given object to a native Go value.
func (doc *Document) Interface() interface{} {
	return doc.PDFDoc
}

// Returns True if the given object is equal to this object.
func (doc *Document) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (doc *Document) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "finish":
		return object.NewBuiltin("document.finish", doc.finish), true
	case "new_page":
		return object.NewBuiltin("document.new_page", doc.newPage), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (doc *Document) SetAttr(name string, value object.Object) error {
	switch name {
	case "language":
		if value.Type() == object.STRING {
			l, err := frontend.GetLanguage(value.(*object.String).Value())
			if err != nil {
				return err
			}
			doc.PDFDoc.SetDefaultLanguage(l)
			return nil
		}
		if value.Type() == "backend.lang" {
			l := value.(*rlang.Lang).Value
			doc.PDFDoc.SetDefaultLanguage(l)
			return nil
		}
	case "title":
		if value.Type() == object.STRING {
			doc.PDFDoc.Title = value.(*object.String).Value()
			return nil
		}
		return object.Errorf("title must be a string")
	}
	return object.Errorf("cannot set attribute %s on document", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (doc *Document) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (doc *Document) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on document", opType)
}

// Cost returns the incremental processing cost of this object.
func (doc *Document) Cost() int {
	return 0
}
