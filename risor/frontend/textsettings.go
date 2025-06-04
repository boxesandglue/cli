package frontend

import (
	"context"
	"sort"

	"github.com/boxesandglue/boxesandglue/backend/bag"
	"github.com/boxesandglue/boxesandglue/frontend"
	rcolor "github.com/boxesandglue/cli/risor/backend/color"
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
	case "backgroundcolor":
		return frontend.SettingBackgroundColor
	case "borderbottomcolor":
		return frontend.SettingBorderBottomColor
	case "borderbottomleftradius":
		return frontend.SettingBorderBottomLeftRadius
	case "borderbottomrightradius":
		return frontend.SettingBorderBottomRightRadius
	case "borderbottomstyle":
		return frontend.SettingBorderBottomStyle
	case "borderbottomwidth":
		return frontend.SettingBorderBottomWidth
	case "borderleftcolor":
		return frontend.SettingBorderLeftColor
	case "borderleftstyle":
		return frontend.SettingBorderLeftStyle
	case "borderleftwidth":
		return frontend.SettingBorderLeftWidth
	case "borderrightcolor":
		return frontend.SettingBorderRightColor
	case "borderrightstyle":
		return frontend.SettingBorderRightStyle
	case "borderrightwidth":
		return frontend.SettingBorderRightWidth
	case "bordertopcolor":
		return frontend.SettingBorderTopColor
	case "bordertopleftradius":
		return frontend.SettingBorderTopLeftRadius
	case "bordertoprightradius":
		return frontend.SettingBorderTopRightRadius
	case "bordertopstyle":
		return frontend.SettingBorderTopStyle
	case "bordertopwidth":
		return frontend.SettingBorderTopWidth
	case "box":
		return frontend.SettingBox
	case "color":
		return frontend.SettingColor
	case "debug":
		return frontend.SettingDebug
	case "fontexpansion":
		return frontend.SettingFontExpansion
	case "fontfamily":
		return frontend.SettingFontFamily
	case "fontweight":
		return frontend.SettingFontWeight
	case "halign":
		return frontend.SettingHAlign
	case "hangingpunctuation":
		return frontend.SettingHangingPunctuation
	case "height":
		return frontend.SettingHeight
	case "hyperlink":
		return frontend.SettingHyperlink
	case "indentleft":
		return frontend.SettingIndentLeft
	case "indentleftrows":
		return frontend.SettingIndentLeftRows
	case "leading":
		return frontend.SettingLeading
	case "marginbottom":
		return frontend.SettingMarginBottom
	case "marginleft":
		return frontend.SettingMarginLeft
	case "marginright":
		return frontend.SettingMarginRight
	case "margintop":
		return frontend.SettingMarginTop
	case "opentypefeature":
		return frontend.SettingOpenTypeFeature
	case "paddingbottom":
		return frontend.SettingPaddingBottom
	case "paddingleft":
		return frontend.SettingPaddingLeft
	case "paddingright":
		return frontend.SettingPaddingRight
	case "paddingtop":
		return frontend.SettingPaddingTop
	case "prepend":
		return frontend.SettingPrepend
	case "preservewhitespace":
		return frontend.SettingPreserveWhitespace
	case "size":
		return frontend.SettingSize
	case "style":
		return frontend.SettingStyle
	case "tabsize":
		return frontend.SettingTabSize
	case "tabsizespaces":
		return frontend.SettingTabSizeSpaces
	case "textdecorationline":
		return frontend.SettingTextDecorationLine
	case "valign":
		return frontend.SettingVAlign
	case "width":
		return frontend.SettingWidth
	case "yoffset":
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

	return object.NewInt(int64(value.(frontend.SettingType)))
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
	case "backgroundcolor":
		if value.Type() == rcolor.BackendColorType {
			v := value.(*rcolor.RColor).Value
			m.txt.Settings[frontend.SettingBackgroundColor] = v
		}
	case "color":
		if value.Type() == rcolor.BackendColorType {
			v := value.(*rcolor.RColor).Value
			m.txt.Settings[frontend.SettingColor] = v
		}
	case "fontfamily":
		m.txt.Settings[frontend.SettingFontFamily] = value.(*FontFamily).Value
	case "halign":
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
	case "fontexpansion":
		switch value.Type() {
		case object.FLOAT:
			// Convert the float value to a font expansion
			m.txt.Settings[frontend.SettingFontExpansion] = value.(*object.Float).Value()
		}
	case "fontstyle":
		switch value.Type() {
		case object.STRING:
			// Convert the string value to a font style
			switch value.(*object.String).String() {
			case "normal":
				m.txt.Settings[frontend.SettingStyle] = frontend.FontStyleNormal
			case "italic":
				m.txt.Settings[frontend.SettingStyle] = frontend.FontStyleItalic
			case "oblique":
				m.txt.Settings[frontend.SettingStyle] = frontend.FontStyleOblique
			default:
				return object.TypeErrorf("type error: invalid value for SettingStyle: %s", value.(*object.String).String())
			}
		}
	case "fontweight":
		switch value.Type() {
		case object.INT:
			// Convert the int value to a font weight
			m.txt.Settings[frontend.SettingFontWeight] = int(value.(*object.Int).Value())
		case object.STRING:
			switch value.(*object.String).String() {
			case "normal":
				m.txt.Settings[frontend.SettingFontWeight] = frontend.FontWeight400
			case "bold":
				m.txt.Settings[frontend.SettingFontWeight] = frontend.FontWeight700
			}
		}
	case "hangingpunctuation":
		switch value.Type() {
		case object.BOOL:
			// Convert the bool value to a hanging punctuation setting
			hp := value.(*object.Bool).Value()
			if hp {
				m.txt.Settings[frontend.SettingHangingPunctuation] = frontend.HangingPunctuation(1)
			}
		}
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
