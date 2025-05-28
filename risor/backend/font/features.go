package font

import (
	"context"

	"github.com/boxesandglue/textlayout/harfbuzz"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const FeatureType = "font.feature"

type Feature struct {
	value harfbuzz.Feature
}

func newFeature(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("font.feature() takes exactly one argument (string)")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("font.feature() expects a string argument")
	}
	featureStr := args[0].(*object.String).Value()
	f, err := harfbuzz.ParseFeature(featureStr)
	if err != nil {
		return object.Errorf("font.feature() failed to parse feature: %w", err)
	}
	return &Feature{value: f}
}

// Type of the object.
func (f *Feature) Type() object.Type {
	return FeatureType
}

// Inspect returns a string representation of the given object.
func (f *Feature) Inspect() string {
	return f.value.Tag.String()
}

// Interface converts the given object to a native Go value.
func (f *Feature) Interface() interface{} {
	return f.value
}

// Returns True if the given object is equal to this object.
func (f *Feature) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (f *Feature) GetAttr(name string) (object.Object, bool) {
	return nil, false // No attributes defined for Feature
}

// SetAttr sets the attribute with the given name on this object.
func (f *Feature) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on feature", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (f *Feature) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (f *Feature) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("font.feature does not support operations: %s", opType)
}

// Cost returns the incremental processing cost of this object.
func (f *Feature) Cost() int {
	return 0
}
