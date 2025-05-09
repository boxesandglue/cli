package frontend

import (
	"context"
	"sort"

	"github.com/boxesandglue/boxesandglue/backend/bag"
	"github.com/boxesandglue/boxesandglue/frontend"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const settingsType = "text.settings"

type settings struct {
	txt *frontend.Text
}

func settingTostring(s frontend.SettingType) string {
	return s.String()
}

func stringToSetting(s string) frontend.SettingType {
	switch s {
	case "SettingBackgroundColor":
		return frontend.SettingBackgroundColor
	case "SettingBorderBottomColor":
		return frontend.SettingBorderBottomColor
	case "SettingBorderBottomLeftRadius":
		return frontend.SettingBorderBottomLeftRadius
	case "SettingBorderBottomRightRadius":
		return frontend.SettingBorderBottomRightRadius
	case "SettingBorderBottomStyle":
		return frontend.SettingBorderBottomStyle
	case "SettingBorderBottomWidth":
		return frontend.SettingBorderBottomWidth
	case "SettingBorderLeftColor":
		return frontend.SettingBorderLeftColor
	case "SettingBorderLeftStyle":
		return frontend.SettingBorderLeftStyle
	case "SettingBorderLeftWidth":
		return frontend.SettingBorderLeftWidth
	case "SettingBorderRightColor":
		return frontend.SettingBorderRightColor
	case "SettingBorderRightStyle":
		return frontend.SettingBorderRightStyle
	case "SettingBorderRightWidth":
		return frontend.SettingBorderRightWidth
	case "SettingBorderTopColor":
		return frontend.SettingBorderTopColor
	case "SettingBorderTopLeftRadius":
		return frontend.SettingBorderTopLeftRadius
	case "SettingBorderTopRightRadius":
		return frontend.SettingBorderTopRightRadius
	case "SettingBorderTopStyle":
		return frontend.SettingBorderTopStyle
	case "SettingBorderTopWidth":
		return frontend.SettingBorderTopWidth
	case "SettingBox":
		return frontend.SettingBox
	case "SettingColor":
		return frontend.SettingColor
	case "SettingDebug":
		return frontend.SettingDebug
	case "SettingFontExpansion":
		return frontend.SettingFontExpansion
	case "SettingFontFamily":
		return frontend.SettingFontFamily
	case "SettingFontWeight":
		return frontend.SettingFontWeight
	case "SettingHAlign":
		return frontend.SettingHAlign
	case "SettingHangingPunctuation":
		return frontend.SettingHangingPunctuation
	case "SettingHeight":
		return frontend.SettingHeight
	case "SettingHyperlink":
		return frontend.SettingHyperlink
	case "SettingIndentLeft":
		return frontend.SettingIndentLeft
	case "SettingIndentLeftRows":
		return frontend.SettingIndentLeftRows
	case "SettingLeading":
		return frontend.SettingLeading
	case "SettingMarginBottom":
		return frontend.SettingMarginBottom
	case "SettingMarginLeft":
		return frontend.SettingMarginLeft
	case "SettingMarginRight":
		return frontend.SettingMarginRight
	case "SettingMarginTop":
		return frontend.SettingMarginTop
	case "SettingOpenTypeFeature":
		return frontend.SettingOpenTypeFeature
	case "SettingPaddingBottom":
		return frontend.SettingPaddingBottom
	case "SettingPaddingLeft":
		return frontend.SettingPaddingLeft
	case "SettingPaddingRight":
		return frontend.SettingPaddingRight
	case "SettingPaddingTop":
		return frontend.SettingPaddingTop
	case "SettingPrepend":
		return frontend.SettingPrepend
	case "SettingPreserveWhitespace":
		return frontend.SettingPreserveWhitespace
	case "SettingSize":
		return frontend.SettingSize
	case "SettingStyle":
		return frontend.SettingStyle
	case "SettingTabSize":
		return frontend.SettingTabSize
	case "SettingTabSizeSpaces":
		return frontend.SettingTabSizeSpaces
	case "SettingTextDecorationLine":
		return frontend.SettingTextDecorationLine
	case "SettingVAlign":
		return frontend.SettingVAlign
	case "SettingWidth":
		return frontend.SettingWidth
	case "SettingYOffset":
		return frontend.SettingYOffset
	}
	return frontend.SettingType(0)
}

func AsSettings(obj object.Object) (*settings, *object.Error) {
	m, ok := obj.(*settings)
	if !ok {
		return nil, object.TypeErrorf("type error: expected a map (%s given)", obj.Type())
	}
	return m, nil
}

func (s *settings) Type() object.Type {
	return settingsType
}

func (s *settings) Inspect() string {
	return "settings"
}

func (m *settings) String() string {
	return m.Inspect()
}

func (m *settings) Value() map[string]object.Object {
	// FIXME: this is a hack to make the settings work with the frontend
	return map[string]object.Object{}
}

func (m *settings) SetAttr(name string, value object.Object) error {
	bag.Logger.Error("Unhandled function", "function", "SetAttr", "where", "textsettings.go")
	return nil
}

func (m *settings) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "keys":
		return object.NewBuiltin("map.keys", func(ctx context.Context, args ...object.Object) object.Object {
			if len(args) != 0 {
				return object.NewArgsError("map.keys", 0, len(args))
			}
			return m.Keys()
		}), true
	case "values":
		return object.NewBuiltin("map.values", func(ctx context.Context, args ...object.Object) object.Object {
			if len(args) != 0 {
				return object.NewArgsError("map.values", 0, len(args))
			}
			return m.Values()
		}), true
	case "get":
		return object.NewBuiltin("map.get", func(ctx context.Context, args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return object.NewArgsRangeError("map.get", 1, 2, len(args))
			}
			return object.NewString("settings/get")
		}), true
	case "clear":
		return object.NewBuiltin("map.clear", func(ctx context.Context, args ...object.Object) object.Object {
			if len(args) != 0 {
				return object.NewArgsError("map.clear", 0, len(args))
			}
			m.Clear()
			return m
		}), true
	case "copy":
		return object.NewBuiltin("map.copy", func(ctx context.Context, args ...object.Object) object.Object {
			if len(args) != 0 {
				return object.NewArgsError("map.copy", 0, len(args))
			}
			return m.Copy()
		}), true
	case "items":
		return object.NewBuiltin("map.items", func(ctx context.Context, args ...object.Object) object.Object {
			if len(args) != 0 {
				return object.NewArgsError("map.items", 0, len(args))
			}
			return m.ListItems()
		}), true
	case "pop":
		return object.NewBuiltin("map.pop", func(ctx context.Context, args ...object.Object) object.Object {
			nArgs := len(args)
			if nArgs < 1 || nArgs > 2 {
				return object.NewArgsRangeError("map.pop", 1, 2, len(args))
			}
			key, err := object.AsString(args[0])
			if err != nil {
				return err
			}
			var def object.Object
			if nArgs == 2 {
				def = args[1]
			}
			return m.Pop(key, def)
		}), true
	case "setdefault":
		return object.NewBuiltin("map.setdefault", func(ctx context.Context, args ...object.Object) object.Object {
			if len(args) != 2 {
				return object.NewArgsError("map.setdefault", 2, len(args))
			}
			key, err := object.AsString(args[0])
			if err != nil {
				return err
			}
			return m.SetDefault(key, args[1])
		}), true
	case "update":
		return object.NewBuiltin("map.update", func(ctx context.Context, args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewArgsError("map.update", 1, len(args))
			}
			other, err := AsSettings(args[0])
			if err != nil {
				return err
			}
			m.Update(other)
			return m
		}), true
	}
	return object.NewString("settings/default get"), true
	// o, ok := m.items[name]
	// return o, ok
}

func (m *settings) ListItems() *object.List {
	bag.Logger.Error("Unhandled function", "function", "ListItems", "where", "textsettings.go")
	return object.NewList(nil)
}

func (s *settings) Clear() {
	s.txt.Settings = frontend.TypesettingSettings{}
}

func (m *settings) Copy() *settings {
	bag.Logger.Error("Unhandled function", "function", "Copy", "where", "textsettings.go")
	return &settings{}
}

func (m *settings) Pop(key string, def object.Object) object.Object {
	bag.Logger.Error("Unhandled function", "function", "Pop", "where", "textsettings.go")
	return object.Nil
}

func (m *settings) SetDefault(key string, value object.Object) object.Object {
	bag.Logger.Error("Unhandled function", "function", "SetDefault", "where", "textsettings.go")
	return object.Nil
}

func (m *settings) Update(other *settings) {
	bag.Logger.Error("Unhandled function", "function", "Update", "where", "textsettings.go")
}

func (s *settings) SortedKeys() []string {
	keys := make([]string, 0, len(s.txt.Items))
	for k := range s.txt.Settings {
		keys = append(keys, settingTostring(k))
	}
	sort.Strings(keys)
	return keys
}

func (s *settings) Keys() *object.List {
	items := make([]object.Object, 0, len(s.txt.Settings))
	for _, k := range s.SortedKeys() {
		items = append(items, object.NewString(k))
	}
	l := object.NewList(items)
	return l
}

func (s *settings) Values() *object.List {
	// bag.Logger.Error("Unhandled function", "function", "Values","where", "textsettings.go")
	l := object.NewList(nil)
	return l
}

func (m *settings) GetWithObject(key *object.String) object.Object {
	value, found := m.txt.Settings[stringToSetting(key.String())]
	if !found {
		return object.Nil
	}
	_ = value
	return object.Nil
}

func (m *settings) Get(key string) object.Object {
	bag.Logger.Error("Unhandled function", "function", "Get", "where", "textsettings.go")
	return object.Nil
}

func (m *settings) GetWithDefault(key string, defaultValue object.Object) object.Object {
	bag.Logger.Error("Unhandled function", "function", "GetWithDefault", "where", "textsettings.go")
	return nil
}

func (s *settings) Delete(key string) object.Object {
	return object.Nil
}

func (s *settings) Interface() any {
	result := make(map[string]any, len(s.txt.Settings))
	for k, v := range s.txt.Settings {
		result[settingTostring(k)] = v
	}
	return result
}

func (m *settings) Equals(other object.Object) object.Object {
	return object.False
}

func (m *settings) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.TypeErrorf("type error: unsupported operation for map: %v", opType)
}

func (m *settings) GetItem(key object.Object) (object.Object, *object.Error) {
	bag.Logger.Error("Unhandled function", "function", "GetItem", "where", "textsettings.go")
	return object.Nil, nil
}

// GetSlice implements the [start:stop] operator for a container type.
func (m *settings) GetSlice(s object.Slice) (object.Object, *object.Error) {
	return nil, object.TypeErrorf("map does not support slice operations")
}

// SetItem assigns a value to the given key in the map.
func (m *settings) SetItem(key, value object.Object) *object.Error {
	strObj, ok := key.(*object.String)
	if !ok {
		return object.TypeErrorf("type error: map key must be a string (got %s)", key.Type())
	}
	switch strObj.String() {
	case "SettingFontfamily":
		m.txt.Settings[frontend.SettingFontFamily] = value.(*FontFamily).Value
	case "SettingHAlign":
		switch value.(*object.String).String() {
		case "left":
			m.txt.Settings[frontend.SettingHAlign] = frontend.HAlignLeft
		case "center":
			m.txt.Settings[frontend.SettingHAlign] = frontend.HAlignCenter
		case "right":
			m.txt.Settings[frontend.SettingHAlign] = frontend.HAlignRight
		case "justify":
			m.txt.Settings[frontend.SettingHAlign] = frontend.HAlignJustified
		default:
			return object.TypeErrorf("type error: invalid value for SettingHAlign: %T", value)
		}
	case "SettingFontWeight":
		m.txt.Settings[frontend.SettingFontWeight] = int(value.(*object.Int).Value())
	default:
		bag.Logger.Error("Unhandled setting", "value", strObj.String())
	}
	return nil
}

// DelItem deletes the item with the given key from the map.
func (m *settings) DelItem(key object.Object) *object.Error {
	bag.Logger.Error("Unhandled function", "function", "DelItem", "where", "textsettings.go")
	return nil
}

// Contains returns true if the given item is found in this container.
func (m *settings) Contains(key object.Object) *object.Bool {
	bag.Logger.Error("Unhandled function", "function", "Contains", "where", "textsettings.go")
	return object.False
}

func (m *settings) IsTruthy() bool {
	return len(m.txt.Settings) > 0
}

// Len returns the number of items in this container.
func (m *settings) Len() *object.Int {
	return object.NewInt(int64(len(m.txt.Settings)))
}

func (m *settings) Iter() object.Iterator {
	return nil
}

func (m *settings) Cost() int {
	return 0
}
