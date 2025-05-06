package document

import (
	"context"

	"github.com/boxesandglue/boxesandglue/backend/bag"
	"github.com/boxesandglue/boxesandglue/backend/document"
	"github.com/boxesandglue/boxesandglue/backend/node"

	backenddoc "github.com/boxesandglue/boxesandglue/backend/document"
	rbag "github.com/boxesandglue/cli/risor/backend/bag"
	rnode "github.com/boxesandglue/cli/risor/backend/node"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

// Page represents a PDF page object.
// It is a wrapper around the backend document page object.
type Page struct {
	Value *document.Page
}

func newPage(d *backenddoc.PDFDocument) *Page {
	return &Page{
		Value: d.NewPage(),
	}
}

func (p *Page) outputAt(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.ArgsErrorf("page.output_at() takes exactly 3 arguments")
	}
	firstArg := args[0]
	secondarg := args[1]
	thirdArg := args[2]
	if firstArg.Type() != "bag.scaledpoint" {
		return object.ArgsErrorf("page.output_at() expects a bag.scaledpoint argument (x-coordinate), got %s", firstArg.Type())
	}
	if secondarg.Type() != "bag.scaledpoint" {
		return object.ArgsErrorf("page.output_at() expects a bag.scaledpoint argument (y-coordinate), got %s", secondarg.Type())
	}
	var n *rnode.Node
	var ok bool
	if n, ok = thirdArg.(*rnode.Node); !ok {
		return object.ArgsErrorf("page.output_at() expects a node.node argument (node)")
	}

	vl := n.Interface().(*node.VList)
	p.Value.OutputAt(firstArg.Interface().(bag.ScaledPoint), secondarg.Interface().(bag.ScaledPoint), vl)
	return nil
}

func (p *Page) shipout(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.ArgsErrorf("page.shipout() takes no arguments")
	}
	p.Value.Shipout()
	return nil
}

// Type of the object.
func (p *Page) Type() object.Type {
	return "backend.document.page"
}

// Inspect returns a string representation of the given object.
func (p *Page) Inspect() string {
	return "Page"
}

// Interface converts the given object to a native Go value.
func (p *Page) Interface() interface{} {
	return p.Value
}

// Equals returns True if the given object is equal to this object.
func (p *Page) Equals(other object.Object) object.Object {
	return object.False
}

/*
	ExtraOffset       bag.ScaledPoint
	Background        []Object
	Objects           []Object
	Userdata          map[any]any
	Finished          bool
	StructureElements []*StructureElement
	Annotations       []pdf.Annotation
	Spotcolors        []*color.Color
	Objectnumber      pdf.Objectnumber
*/

// GetAttr returns the attribute with the given name from this object.
func (p *Page) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "height":
		return &rbag.RSP{Value: p.Value.Height}, true
	case "output_at":
		return object.NewBuiltin("page.output_at", p.outputAt), true
	case "shipout":
		return object.NewBuiltin("page.shipout", p.shipout), true
	case "width":
		return &rbag.RSP{Value: p.Value.Width}, true
	}

	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (p *Page) SetAttr(name string, value object.Object) error {
	switch name {
	case "width":
		if value.Type() != rbag.ScaledPointType {
			return object.ArgsErrorf("page.width expects a bag.scaledpoint argument (width)")
		}
		p.Value.Width = value.(*rbag.RSP).Value
		return nil
	case "height":
		if value.Type() != rbag.ScaledPointType {
			return object.ArgsErrorf("page.height expects a bag.scaledpoint argument (height)")
		}
		p.Value.Height = value.(*rbag.RSP).Value
		return nil
	}
	return object.Errorf("cannot set attribute %s on page", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (p *Page) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (p *Page) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on page", opType)
}

// Cost returns the incremental processing cost of this object.
func (p *Page) Cost() int {
	return 0
}
