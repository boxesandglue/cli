package bag

import (
	"context"
	"fmt"

	"github.com/boxesandglue/boxesandglue/backend/bag"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type RSP struct {
	Value bag.ScaledPoint
}

// Type of the object.
func (sp *RSP) Type() object.Type {
	return "backend.sp"
}

// Inspect returns a string representation of the given object.
func (sp *RSP) Inspect() string {
	return sp.Value.String()
}

// Interface converts the given object to a native Go value.
func (sp *RSP) Interface() interface{} {
	return sp.Value
}

// Returns True if the given object is equal to this object.
func (sp *RSP) Equals(other object.Object) object.Object {
	if other.Type() != sp.Type() {
		return object.False
	}
	return object.NewBool(sp.Value == other.(*RSP).Value)
}

// GetAttr returns the attribute with the given name from this object.
func (sp *RSP) GetAttr(name string) (object.Object, bool) {
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (sp *RSP) SetAttr(name string, value object.Object) error {
	return fmt.Errorf("cannot set attribute %s on SP", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (sp *RSP) IsTruthy() bool {
	return sp.Value != 0
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (sp *RSP) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	if opType == op.Multiply {
		if right.Type() == object.INT {
			return &RSP{Value: bag.MultiplyFloat(sp.Value, float64(right.Interface().(int64)))}
		}
	}
	return object.Errorf("operation %s not supported on SP", opType)
}

// Cost returns the incremental processing cost of this object.
func (sp *RSP) Cost() int {
	return 0
}

func bagSP(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("bag.sp() expects one argument")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("bag.sp() expects a string argument (a unit)")
	}

	firstArg := args[0].(*object.String).Value()
	sp, err := bag.SP(firstArg)
	if err != nil {
		return object.Errorf("bag.sp() failed: %s", err)
	}
	return &RSP{Value: sp}
}

// Module returns the bag module.
func Module() *object.Module {
	return object.NewBuiltinsModule("bag", map[string]object.Object{
		"sp": object.NewBuiltin("bag.sp", bagSP),
	})
}
