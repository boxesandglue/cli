package frontend

import (
	"context"

	"github.com/boxesandglue/boxesandglue/frontend"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type FontFamily struct {
	Value *frontend.FontFamily
}

func (ff *FontFamily) addMember(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("frontend.add_member() takes exactly one argument")
	}
	firstArg := args[0]
	if firstArg.Type() != object.MAP {
		return object.ArgsErrorf("frontend.add_member() expects a map argument (font source, weight and style)")
	}
	mapArg := firstArg.(*object.Map)
	var weight frontend.FontWeight
	var style frontend.FontStyle
	var fs *frontend.FontSource
	if value := mapArg.Get("weight"); value != object.Nil {
		if value.Type() != object.INT {
			return object.ArgsErrorf("frontend.add_member() expects an int argument (weight)")
		}
		weight = frontend.FontWeight(value.(*object.Int).Value())
	}
	if value := mapArg.Get("style"); value != object.Nil {
		if value.Type() != object.STRING {
			return object.ArgsErrorf("frontend.add_member() expects a string argument (style)")
		}
		styleString := value.(*object.String).Value()
		switch styleString {
		case "normal":
			style = frontend.FontStyleNormal
		case "italic":
			style = frontend.FontStyleItalic
		case "oblique":
			style = frontend.FontStyleOblique
		default:
			return object.ArgsErrorf("frontend.add_member() expects a string argument (style) with value normal, italic or oblique")
		}
	}
	if value := mapArg.Get("source"); value != object.Nil {
		if value.Type() != "frontend.fontsource" {
			return object.ArgsErrorf("frontend.add_member() expects a font source argument (font source)")
		}
		risorFS := value.(*fontSource)
		fs = &frontend.FontSource{Location: risorFS.location, Name: risorFS.name, Index: risorFS.index}
	}
	err := ff.Value.AddMember(fs, weight, style)
	if err != nil {
		return object.Errorf("frontend.add_member() failed: %s", err)
	}
	return object.Nil
}

// Type of the object.
func (ff *FontFamily) Type() object.Type {
	return "frontend.fontfamily"
}

// Inspect returns a string representation of the given object.
func (ff *FontFamily) Inspect() string {
	return ff.Value.Name
}

// Interface converts the given object to a native Go value.
func (ff *FontFamily) Interface() interface{} {
	return ff.Value
}

// Returns True if the given object is equal to this object.
func (ff *FontFamily) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (ff *FontFamily) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "add_member":
		return object.NewBuiltin("frontend.add_member", ff.addMember), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (ff *FontFamily) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on font family", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (ff *FontFamily) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (ff *FontFamily) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on font family", opType)
}

// Cost returns the incremental processing cost of this object.
func (ff *FontFamily) Cost() int {
	return 0
}
