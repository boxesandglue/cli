package pdf

import (
	"context"
	"fmt"

	pdf "github.com/boxesandglue/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

// TypeFace is the type name for PDFFace objects.
const TypeFace = "pdf.face"

// Face represents a PDF face object in Risor.
type Face struct {
	Value *pdf.Face
}

// Type of the object.
func (face *Face) Type() object.Type {
	return TypeFace
}

// Inspect returns a string representation of the given object.
func (face *Face) Inspect() string {
	return face.Value.Filename
}

// Interface converts the given object to a native Go value.
func (face *Face) Interface() any {
	return face.Value
}

// Equals returns True if the given object is equal to this object.
func (face *Face) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (face *Face) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "codepoint":
		return object.NewBuiltin("pdf.face.codepoint", face.codepoint), true
	case "codepoints":
		return object.NewBuiltin("pdf.face.codepoints", face.codepoints), true
	case "face_id":
		return object.NewInt(int64(face.Value.FaceID)), true
	case "filename":
		return object.NewString(face.Value.Filename), true
	case "internal_name":
		return object.NewString(face.Value.InternalName()), true
	case "postscript_name":
		return object.NewString(face.Value.PostscriptName), true
	case "register_codepoint":
		return object.NewBuiltin("pdf.face.register_codepoint", face.registerCodepoint), true
	case "register_codepoints":
		return object.NewBuiltin("pdf.face.register_codepoints", face.registerCodepoints), true
	case "units_per_em":
		return object.NewInt(int64(face.Value.UnitsPerEM)), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (face *Face) SetAttr(name string, value object.Object) error {
	return object.Errorf("cannot set attribute %s on text", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (face *Face) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (face *Face) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on Face", opType)
}

// Cost returns the incremental processing cost of this object.
func (face *Face) Cost() int {
	return 0
}

func (face *Face) codepoint(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("pdf.face.codepoint() takes exactly one argument")
	}
	runeObj := args[0]
	if runeObj.Type() != object.INT {
		return object.Errorf("expected int for rune, got %s", runeObj.Type())
	}
	r := rune(runeObj.(*object.Int).Value())
	codepoint := face.Value.Codepoint(r)
	return object.NewInt(int64(codepoint))
}

func (face *Face) codepoints(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("pdf.face.codepoints() takes exactly one argument")
	}
	runesObj := args[0]
	if runesObj.Type() != object.LIST {
		return object.Errorf("expected array for runes, got %s", runesObj.Type())
	}
	runesArray := runesObj.(*object.List)
	runes := make([]rune, 0, runesArray.Len().Value())
	for _, elem := range runesArray.Value() {
		if elem.Type() != object.INT {
			return object.Errorf("expected int in runes array, got %s", elem.Type())
		}
		r := rune(elem.(*object.Int).Value())
		runes = append(runes, r)
	}
	fmt.Println(runes)
	codepoints := face.Value.Codepoints(runes)
	codepointsList := object.NewList(nil)

	for _, cp := range codepoints {
		codepointsList.Append(object.NewInt(int64(cp)))
	}
	return codepointsList
}

func (face *Face) registerCodepoint(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("pdf.face.register_codepoint() takes exactly one argument")
	}
	codepointObj := args[0]
	if codepointObj.Type() != object.INT {
		return object.Errorf("expected int for codepoint, got %s", codepointObj.Type())
	}
	codepoint := int(codepointObj.(*object.Int).Value())
	face.Value.RegisterCodepoint(codepoint)
	return object.Nil
}

func (face *Face) registerCodepoints(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("pdf.face.register_codepoints() takes exactly one argument")
	}
	codepointsObj := args[0]
	if codepointsObj.Type() != object.LIST {
		return object.Errorf("expected array for codepoints, got %s", codepointsObj.Type())
	}
	codepointsArray := codepointsObj.(*object.List)
	codepoints := make([]int, 0, codepointsArray.Len().Value())
	for _, elem := range codepointsArray.Value() {
		if elem.Type() != object.INT {
			return object.Errorf("expected int in codepoints array, got %s", elem.Type())
		}
		cp := int(elem.(*object.Int).Value())
		codepoints = append(codepoints, cp)

	}
	face.Value.RegisterCodepoints(codepoints)
	return object.Nil
}
