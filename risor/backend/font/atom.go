package font

import (
	"context"
	"fmt"
	"strings"

	"github.com/boxesandglue/boxesandglue/backend/font"
	"github.com/boxesandglue/cli/risor/backend/bag"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const AtomType = "font.atom"
const AtomsType = "font.atoms"

func newAtom(ctx context.Context, args ...object.Object) object.Object {
	if len(args) == 1 {
		firstArg := args[0]
		if firstArg.Type() != object.STRING {
			return object.ArgsErrorf("font.atom() expects a string argument")
		}
		return &RAtom{Value: &font.Atom{Components: firstArg.Inspect()}}

	}
	return &RAtom{Value: &font.Atom{}}
}

type RAtom struct {
	Value *font.Atom
	pos   int // current position in the atoms slice
}

// Type of the object.
func (ra *RAtom) Type() object.Type {
	return AtomType
}

// Inspect returns a string representation of the given object.
func (ra *RAtom) Inspect() string {
	var sb strings.Builder
	sb.WriteString("font.atom(")
	sb.WriteString(ra.Value.Components)
	sb.WriteString(")")
	if ra.Value.Advance != 0 {
		sb.WriteString(" advance=")
		sb.WriteString(ra.Value.Advance.String())
	}
	if ra.Value.Depth != 0 {
		sb.WriteString(" depth=")
		sb.WriteString(ra.Value.Depth.String())
	}
	if ra.Value.Height != 0 {
		sb.WriteString(" height=")
		sb.WriteString(ra.Value.Height.String())
	}
	if ra.Value.Kernafter != 0 {
		sb.WriteString(" kernafter=")
		sb.WriteString(ra.Value.Kernafter.String())
	}
	if ra.Value.Codepoint != 0 {
		sb.WriteString(" codepoint=")
		sb.WriteString(fmt.Sprintf("%d", ra.Value.Codepoint))
	}
	if ra.Value.IsSpace {
		sb.WriteString(" is_space=true")
	}
	if ra.Value.Hyphenate {
		sb.WriteString(" hyphenate=true")
	}
	return sb.String()
}

// Interface converts the given object to a native Go value.
func (ra *RAtom) Interface() interface{} {
	return ra.Value
}

// Returns True if the given object is equal to this object.
func (ra *RAtom) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (ra *RAtom) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "advance":
		return &bag.RSP{Value: ra.Value.Advance}, true
	case "components":
		return object.NewString(ra.Value.Components), true
	case "codepoint":
		return object.NewInt(int64(ra.Value.Codepoint)), true
	case "depth":
		return &bag.RSP{Value: ra.Value.Depth}, true
	case "is_space":
		return object.NewBool(ra.Value.IsSpace), true
	case "height":
		return &bag.RSP{Value: ra.Value.Height}, true
	case "hyphenate":
		return object.NewBool(ra.Value.Hyphenate), true
	case "kernafter":
		return &bag.RSP{Value: ra.Value.Kernafter}, true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (ra *RAtom) SetAttr(name string, value object.Object) error {
	panic("not implemented") // TODO: Implement
}

// IsTruthy returns true if the object is considered "truthy".
func (ra *RAtom) IsTruthy() bool {
	panic("not implemented") // TODO: Implement
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (ra *RAtom) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	panic("not implemented") // TODO: Implement
}

// Cost returns the incremental processing cost of this object.
func (ra *RAtom) Cost() int {
	panic("not implemented") // TODO: Implement
}

type RAtoms struct {
	Value   []font.Atom
	current *RAtom
	pos     int
}

// Type of the object.
func (a *RAtoms) Type() object.Type {
	return AtomType
}

// Inspect returns a string representation of the given object.
func (a *RAtoms) Inspect() string {
	return "font.atom"
}

// Interface converts the given object to a native Go value.
func (a *RAtoms) Interface() interface{} {
	return a.Value
}

// Returns True if the given object is equal to this object.
func (a *RAtoms) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (a *RAtoms) GetAttr(name string) (object.Object, bool) {
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (a *RAtoms) SetAttr(name string, value object.Object) error {
	return object.ArgsErrorf("font.atom does not support setting attributes")
}

// IsTruthy returns true if the object is considered "truthy".
func (a *RAtoms) IsTruthy() bool {
	return len(a.Value) > 0
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (a *RAtoms) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("font.atom does not support operations: %s", opType)
}

// Cost returns the incremental processing cost of this object.
func (a *RAtoms) Cost() int {
	return 0
}

// Next advances the iterator and then returns the current object and a
// bool indicating whether the returned item is valid. Once Next() has been
// called, the Entry() method can be used to get an IteratorEntry.
func (a *RAtoms) Next(_ context.Context) (object.Object, bool) {
	if a.pos >= len(a.Value) {
		return nil, false
	}

	a.current = &RAtom{Value: &a.Value[a.pos], pos: a.pos}
	a.pos++
	return a.current, true

}

// Entry returns the current entry in the iterator and a bool indicating
// whether the returned item is valid.
func (a *RAtoms) Entry() (object.IteratorEntry, bool) {
	if a.current == nil {
		return nil, false
	}
	return object.NewEntry(object.NewInt(int64(a.pos-1)), a.current), true
}
