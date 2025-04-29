package frontend

import (
	"github.com/boxesandglue/boxesandglue/backend/bag"
	"github.com/boxesandglue/boxesandglue/frontend"
	rbag "github.com/boxesandglue/cli/risor/backend/bag"
	rdocument "github.com/boxesandglue/cli/risor/backend/document"
	rnode "github.com/boxesandglue/cli/risor/backend/node"

	"context"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type frontendDocument struct {
	// The filename of the PDF file
	value *frontend.Document
	// The document object
	doc *rdocument.Document
}

func (fd *frontendDocument) newFontFamily(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("frontend.new_fontfamily() takes exactly one argument")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("frontend.new_fontfamily() expects a string argument (font family name)")
	}
	familyName := args[0].(*object.String).Value()
	return &FontFamily{Value: fd.value.NewFontFamily(familyName)}
}

func (fd *frontendDocument) formatParagraph(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("frontend.format_paragraph() takes exactly one argument")
	}
	firstArg := args[0]
	if firstArg.Type() != object.MAP {
		return object.ArgsErrorf("frontend.format_paragraph() expects a map argument (formatting options)")
	}
	lst := firstArg.(*object.Map).Value()
	var opts = make([]frontend.TypesettingOption, 0)
	var wd bag.ScaledPoint
	var ftext *frontend.Text
	for k, v := range lst {
		switch k {
		case "width":
			if v.Type() == "bag.scaledpoint" {
				wd = v.(*rbag.RSP).Value
			} else {
				return object.ArgsErrorf("frontend.format_paragraph() expects a bag.scaledpoint argument (width)")
			}
		case "text":
			if v.Type() == "frontend.text" {
				ftext = v.(*text).Value
			} else {
				return object.ArgsErrorf("frontend.format_paragraph() expects a frontend.text argument (text)")
			}
		case "leading":
			if v.Type() == "bag.scaledpoint" {
				opts = append(opts, frontend.Leading(v.(*rbag.RSP).Value))
			} else {
				return object.ArgsErrorf("frontend.format_paragraph() expects a bag.scaledpoint argument (leading)")
			}
		case "font_size":
			if v.Type() == "bag.scaledpoint" {
				opts = append(opts, frontend.FontSize(v.(*rbag.RSP).Value))
			} else {
				return object.ArgsErrorf("frontend.format_paragraph() expects a bag.scaledpoint argument (font_size)")
			}
		case "family":
			if v.Type() == "frontend.fontfamily" {
				ff := v.(*FontFamily)
				opts = append(opts, frontend.Family(ff.Value))
			} else {
				return object.ArgsErrorf("frontend.format_paragraph() expects a frontend.fontfamily argument (font family)")
			}
		default:
			// fmt.Println(`~~> k,v`, k, v)
		}
	}
	vlist, _, err := fd.value.FormatParagraph(ftext, wd, opts...)
	if err != nil {
		return object.NewError(err)
	}
	vl := &rnode.Node{Value: vlist}
	return vl
}

// Type of the object.
func (fd *frontendDocument) Type() object.Type {
	return "frontend.document"
}

// Inspect returns a string representation of the given object.
func (fd *frontendDocument) Inspect() string {
	return fd.value.Doc.Filename
}

// Interface converts the given object to a native Go value.
func (fd *frontendDocument) Interface() interface{} {
	return fd.value
}

// Returns True if the given object is equal to this object.
func (fd *frontendDocument) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (fd *frontendDocument) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "doc":
		return fd.doc, true
	case "new_fontfamily":
		return object.NewBuiltin("frontend.new_fontfamily", fd.newFontFamily), true
	case "format_paragraph":
		return object.NewBuiltin("frontend.format_paragraph", fd.formatParagraph), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (fd *frontendDocument) SetAttr(name string, value object.Object) error {
	return nil
}

// IsTruthy returns true if the object is considered "truthy".
func (fd *frontendDocument) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (fd *frontendDocument) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on frontendDocument", opType)
}

// Cost returns the incremental processing cost of this object.
func (fd *frontendDocument) Cost() int {
	return 0
}
