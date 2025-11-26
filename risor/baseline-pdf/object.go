package pdf

import (
	"context"
	"fmt"

	pdf "github.com/boxesandglue/baseline-pdf"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

func (obj *Object) pdfObjectSave(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.ArgsErrorf("pdf.object.save() takes no arguments")
	}
	err := obj.Value.Save()
	if err != nil {
		return object.Errorf("failed to save PDF object: %v", err)
	}
	return object.Nil
}

// TypeObject is the type name for PDF object.
const TypeObject = "pdf.object"

// Object represents a PDF face object in Risor.
type Object struct {
	Value *pdf.Object
	rMap  *object.Map
}

// Type of the object.
func (obj *Object) Type() object.Type {
	return TypeObject
}

// Inspect returns a string representation of the given object.
func (obj *Object) Inspect() string {
	return obj.Value.ObjectNumber.String()
}

// Interface converts the given object to a native Go value.
func (obj *Object) Interface() any {
	return obj.Value
}

// Equals returns True if the given object is equal to this object.
func (obj *Object) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (obj *Object) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "array":
		arr := object.NewList(nil)
		for _, v := range obj.Value.Array {
			switch val := v.(type) {
			case string:
				arr.Append(object.NewString(val))
			case int:
				arr.Append(object.NewInt(int64(val)))
			case float64:
				arr.Append(object.NewFloat(val))
			default:
				arr.Append(object.NewString(fmt.Sprintf("%v", val)))
			}
		}
		return arr, true
	case "data":
		return object.NewBuffer(obj.Value.Data), true
	case "dictionary":
		if obj.rMap == nil {
			m := make(map[string]object.Object)
			obj.rMap = object.NewMap(m)
			return obj.rMap, true
		}
		return obj.rMap, true
	case "force_stream":
		return object.NewBool(obj.Value.ForceStream), true
	case "object_number":
		return objectNumber{Value: obj.Value.ObjectNumber}, true
	case "raw":
		return object.NewBool(obj.Value.Raw), true
	case "save":
		var err error
		if obj.rMap != nil && obj.rMap.Size() > 0 {
			obj.Value.Dictionary, err = fillDictFromMap(obj.rMap)
			if err != nil {
				return object.NewError(err), true
			}
		}
		return object.NewBuiltin("pdf.object.save", obj.pdfObjectSave), true
	case "set_compression":
		return object.NewBuiltin("pdf.object.set_compression", func(ctx context.Context, args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.ArgsErrorf("pdf.object.set_compression() takes one argument (int)")
			}
			if args[0].Type() != object.INT {
				return object.ArgsErrorf("pdf.object.set_compression() expects an int argument")
			}
			level := int(args[0].(*object.Int).Value())
			obj.Value.SetCompression(uint(level))
			return object.Nil
		}), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (obj *Object) SetAttr(name string, value object.Object) error {
	switch name {
	case "object_number":
		switch value.Type() {
		case typeObjectNumber:
			obj.Value.ObjectNumber = value.(objectNumber).Value
			return nil
		case object.INT:
			intVal := value.(*object.Int).Value()
			obj.Value.ObjectNumber = pdf.Objectnumber(intVal)
			return nil
		default:
			return object.Errorf("expected baseline-pdf.objectnumber or int for object_number, got %s", value.Type())
		}
	case "dictionary":
		if value.Type() == object.MAP {
			var err error
			obj.Value.Dictionary, err = fillDictFromMap(value.(*object.Map))
			if err != nil {
				return err
			}
			return nil
		}
		return object.Errorf("dictionary must be a map")
	case "array":
		if value.Type() == object.LIST {
			arr := value.(*object.List)
			pdfArr := make([]any, 0, len(arr.Value()))
			for _, v := range arr.Value() {
				switch v.Type() {
				case object.INT:
					pdfArr = append(pdfArr, int(v.(*object.Int).Value()))
				case object.FLOAT:
					pdfArr = append(pdfArr, v.(*object.Float).Value())
				case object.STRING:
					pdfArr = append(pdfArr, v.(*object.String).Value())
				default:
					return object.Errorf("unsupported type %s in array", v.Type())
				}
			}
			obj.Value.Array = pdfArr
			return nil
		}
		return object.Errorf("array must be a list")
	case "data":
		if value.Type() == object.BUFFER {
			obj.Value.Data = value.(*object.Buffer).Value()
			return nil
		}
		return object.Errorf("data must be a buffer")
	case "force_stream":
		if value.Type() == object.BOOL {
			obj.Value.ForceStream = value.(*object.Bool).Value()
			return nil
		}
		return object.Errorf("force_stream must be a bool")
	case "raw":
		if value.Type() == object.BOOL {
			obj.Value.Raw = value.(*object.Bool).Value()
			return nil
		}
		return object.Errorf("raw must be a bool")
	}

	return object.Errorf("cannot set attribute %s on object", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (obj *Object) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (obj *Object) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on Object", opType)
}

// Cost returns the incremental processing cost of this object.
func (obj *Object) Cost() int {
	return 0
}
