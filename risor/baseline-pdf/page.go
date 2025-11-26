package pdf

import (
	"fmt"

	pdf "github.com/boxesandglue/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

// TypePage is the type name for PDFPage objects.
const TypePage = "pdf.page"

// Page represents a PDF page object in Risor.
type Page struct {
	Value *pdf.Page
}

// Type of the object.
func (pg *Page) Type() object.Type {
	return TypePage
}

// Inspect returns a string representation of the given object.
func (pg *Page) Inspect() string {
	return pg.Value.Objnum.String()
}

// Interface converts the given object to a native Go value.
func (pg *Page) Interface() any {
	return pg.Value
}

// Equals returns True if the given object is equal to this object.
func (pg *Page) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (pg *Page) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "faces":
		faces := object.NewList(nil)
		for _, f := range pg.Value.Faces {
			rf := &Face{Value: f}
			faces.Append(rf)
		}
		return faces, true
	case "imagefiles":
		imgfiles := object.NewList(nil)
		for _, imf := range pg.Value.Images {
			rimf := &ImageFile{Value: imf}
			imgfiles.Append(rimf)
		}
		return imgfiles, true
	case "width":
		return object.NewFloat(pg.Value.Width), true
	case "height":
		return object.NewFloat(pg.Value.Height), true
	case "object_number":
		return &objectNumber{Value: pg.Value.Objnum}, true
	case "offset_x":
		return object.NewFloat(pg.Value.OffsetX), true
	case "offset_y":
		return object.NewFloat(pg.Value.OffsetY), true
	case "dict":
		rMap := object.NewMap(nil)
		for k, v := range pg.Value.Dict {
			valueSet := convertPDFToRisorObject(v)
			rMap.Set(string(k), valueSet)
		}
		return rMap, true
	default:
		fmt.Printf("Unknown get attribute %s\n", name)
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (pg *Page) SetAttr(name string, value object.Object) error {
	switch name {
	case "dict":
		if value.Type() == object.MAP {
			var err error
			pg.Value.Dict, err = fillDictFromMap(value.(*object.Map))
			if err != nil {
				return err
			}
			return nil
		}
		return object.Errorf("viewer_preferences must be a map")
	case "faces":
		if value.Type() == object.LIST {
			arr := value.(*object.List)
			faces := make([]*pdf.Face, 0, len(arr.Value()))
			for _, v := range arr.Value() {
				if v.Type() != TypeFace {
					return object.Errorf("expected pdf.face in faces array, got %s", v.Type())
				}
				faces = append(faces, v.(*Face).Value)
			}
			pg.Value.Faces = faces
			return nil
		}
		return object.Errorf("faces must be an array")
	case "height":
		switch value.Type() {
		case object.INT:
			pg.Value.Height = float64(value.(*object.Int).Value())
			return nil
		case object.FLOAT:
			pg.Value.Height = value.(*object.Float).Value()
			return nil
		default:
			return object.Errorf("expected int or float for height, got %s", value.Type())
		}
	case "images":
		if value.Type() == object.LIST {
			arr := value.(*object.List)
			imgfiles := make([]*pdf.Imagefile, 0, len(arr.Value()))
			for _, v := range arr.Value() {
				if v.Type() != "baseline-pdf.imagefile" {
					return object.Errorf("expected baseline-pdf.imagefile in imagefiles array, got %s", v.Type())
				}
				imgfiles = append(imgfiles, v.(*ImageFile).Value)
			}
			pg.Value.Images = imgfiles
			return nil
		}
		return object.Errorf("imagefiles must be an array")
	case "object_number":
		switch value.Type() {
		case typeObjectNumber:
			pg.Value.Objnum = value.(objectNumber).Value
			return nil
		case object.INT:
			intVal := value.(*object.Int).Value()
			pg.Value.Objnum = pdf.Objectnumber(intVal)
			return nil
		default:
			return object.Errorf("expected baseline-pdf.objectnumber or int for object_number, got %s", value.Type())
		}
	case "offset_x":
		switch value.Type() {
		case object.INT:
			pg.Value.OffsetX = float64(value.(*object.Int).Value())
			return nil
		case object.FLOAT:
			pg.Value.OffsetX = value.(*object.Float).Value()
			return nil
		default:
			return object.Errorf("expected int or float for offset_x, got %s", value.Type())
		}
	case "offset_y":
		switch value.Type() {
		case object.INT:
			pg.Value.OffsetY = float64(value.(*object.Int).Value())
			return nil
		case object.FLOAT:
			pg.Value.OffsetY = value.(*object.Float).Value()
			return nil
		default:
			return object.Errorf("expected int or float for offset_y, got %s", value.Type())
		}
	case "width":
		switch value.Type() {
		case object.INT:
			pg.Value.Width = float64(value.(*object.Int).Value())
			return nil
		case object.FLOAT:
			pg.Value.Width = value.(*object.Float).Value()
			return nil
		default:
			return object.Errorf("expected int or float for width, got %s", value.Type())
		}

	}
	return object.Errorf("cannot set attribute %s on page", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (pg *Page) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (pg *Page) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on Face", opType)
}

// Cost returns the incremental processing cost of this object.
func (pg *Page) Cost() int {
	return 0
}
