package font

import (
	"context"

	"github.com/boxesandglue/boxesandglue/backend/font"
	rbag "github.com/boxesandglue/cli/risor/backend/bag"
	rpdf "github.com/boxesandglue/cli/risor/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type Font struct {
	Value *font.Font
}

// newFont expects a pdf.Face and a size argument.
func newFont(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.ArgsErrorf("font.new() takes exactly two arguments")
	}
	if args[0].Type() != rpdf.TypeFace {
		return object.ArgsErrorf("font.new() expects a font argument (pdf.face)")
	}

	if args[1].Type() != rbag.ScaledPointType {
		return object.ArgsErrorf("font.new() expects a size argument (scaledpoint)")
	}
	face := args[0].(*rpdf.Face).Value
	size := args[1].(*rbag.RSP).Value
	fnt := &Font{
		Value: font.NewFont(face, size),
	}

	return fnt
}

// Type of the object.
func (fnt *Font) Type() object.Type {
	return "font.font"
}

// Inspect returns a string representation of the given object.
func (fnt *Font) Inspect() string {
	return fnt.Value.Face.Filename
}

// Interface converts the given object to a native Go value.
func (fnt *Font) Interface() interface{} {
	return fnt.Value
}

// Equals returns True if the given object is equal to this object.
func (fnt *Font) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (fnt *Font) GetAttr(name string) (object.Object, bool) {
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (fnt *Font) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on font", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (fnt *Font) IsTruthy() bool {
	return false
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (fnt *Font) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("cannot run operation %s on font", opType)
}

// Cost returns the incremental processing cost of this object.
func (fnt *Font) Cost() int {
	return 0
}

// Module returns the font module.
func Module() *object.Module {
	return object.NewBuiltinsModule("font", map[string]object.Object{
		"new": object.NewBuiltin("font.new", newFont),
	})
}
