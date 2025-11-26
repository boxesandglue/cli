package pdf

import (
	"fmt"

	pdf "github.com/boxesandglue/baseline-pdf"
	"github.com/risor-io/risor/object"
)

func convertPDFToRisorObject(v any) object.Object {
	switch t := v.(type) {
	case pdf.Dict:
		rMap := object.NewMap(nil)
		for k, v := range t {
			valueSet := convertPDFToRisorObject(v)
			rMap.Set(string(k), valueSet)
		}
		return rMap
	case pdf.String:
		return object.NewString(string(t))
	case int:
		return object.NewInt(int64(t))
	case float64:
		return object.NewFloat(t)
	case []any:
		arr := make([]object.Object, 0, len(t))
		for _, itm := range t {
			arr = append(arr, convertPDFToRisorObject(itm))
		}
		return object.NewList(arr)
	case int64:
		return object.NewInt(t)
	default:
		fmt.Printf("~~> t %T\n", t)
		return object.NewString("<unknown>")
	}
}

func fillDictFromMap(m *object.Map) (pdf.Dict, error) {
	dict := make(pdf.Dict)
	var err error
	keys := m.Keys()
	for _, k := range keys.Value() {
		switch t := k.(type) {
		case *object.String:
			key := t.Value()
			keyPDFName := pdf.Name(key)
			switch v := m.Get(key); v.Type() {
			case object.STRING:
				dict[keyPDFName] = pdf.String(v.(*object.String).Value())
			case object.INT:
				dict[keyPDFName] = v.(*object.Int).Value()
			case object.MAP:
				t := v.(*object.Map)
				dict[keyPDFName], err = fillDictFromMap(t)
				if err != nil {
					return nil, err
				}
			case object.LIST:
				t := v.(*object.List)
				var arr []any
				for _, itm := range t.Value() {
					switch it := itm.(type) {
					case *object.String:
						arr = append(arr, pdf.String(it.Value()))
					case *object.Int:
						arr = append(arr, it.Value())
					default:
						return nil, object.ArgsErrorf("array items must be strings or ints, got %s", it.Type())
					}
				}
				dict[keyPDFName] = arr
			default:
				return nil, object.ArgsErrorf("dictionary values must be strings or ints, got %s", v.Type())
			}
		default:
			// error message if not a string
			return nil, object.ArgsErrorf("dictionary keys must be strings, got %s", t.Type())
		}
	}

	return dict, nil
}
