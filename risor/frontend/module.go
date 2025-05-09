package frontend

import (
	"context"

	"github.com/boxesandglue/boxesandglue/frontend"
	"github.com/boxesandglue/cli/risor/backend/document"
	rlang "github.com/boxesandglue/cli/risor/backend/lang"
	"github.com/risor-io/risor/object"
)

func frontendNewFontsource(ctx context.Context, args ...object.Object) object.Object {
	fs := &fontSource{}
	if len(args) != 1 {
		return object.ArgsErrorf("frontend.new_fontsource() takes exactly one argument")
	}
	if args[0].Type() != object.MAP {
		return object.ArgsErrorf("frontend.new_fontsource() expects a map argument (font source)")
	}
	firstArg := args[0].(*object.Map)

	if value := firstArg.Get("location"); value != object.Nil {
		if value.Type() != object.STRING {
			return object.ArgsErrorf("frontend.new_fontsource() expects a string argument (location)")
		}
		fs.location = value.(*object.String).Value()
	}
	if value := firstArg.Get("name"); value != object.Nil {
		if value.Type() != object.STRING {
			return object.ArgsErrorf("frontend.new_fontsource() expects a string argument (name)")
		}
		fs.name = value.(*object.String).Value()
	}
	// index
	if value := firstArg.Get("index"); value != object.Nil {
		if value.Type() != object.INT {
			return object.ArgsErrorf("frontend.new_fontsource() expects an int argument (index)")
		}
		fs.index = int(value.(*object.Int).Value())
	}

	// fontFeatures
	if value := firstArg.Get("features"); value != object.Nil {
		if value.Type() != object.LIST {
			return object.ArgsErrorf("frontend.new_fontsource() expects a list argument (features)")
		}
		fontFeatures := value.(*object.List).Value()

		for _, ff := range fontFeatures {
			if ff.Type() != object.STRING {
				return object.ArgsErrorf("frontend.new_fontsource() expects a list of strings (fontFeatures)")
			}
			fs.fontFeatures = append(fs.fontFeatures, ff.(*object.String).Value())
		}
	}

	return fs
}

func frontendGetLanguage(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("frontend.get_language() takes exactly one argument")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("frontend.get_language() expects a string argument (language name)")
	}
	l, err := frontend.GetLanguage(args[0].(*object.String).Value())
	if err != nil {
		return object.NewError(err)
	}
	backendLang := &rlang.Lang{Value: l}
	return backendLang
}

func frontendNew(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("frontend.new() takes exactly one argument")
	}

	filename, ok := args[0].(*object.String)
	if !ok {
		return object.ArgsErrorf("frontend.new() expects a string argument (filename of the PDF file)")
	}
	doc, err := frontend.New(filename.Value())
	if err != nil {
		return object.NewError(err)
	}
	fd := &frontendDocument{value: doc, doc: &document.Document{PDFDoc: doc.Doc, Attachments: object.NewList(nil)}}
	return fd
}

func newText(ctx context.Context, args ...object.Object) object.Object {
	return &text{Value: frontend.NewText()}
}

// Module returns the frontend module.
func Module() *object.Module {
	return object.NewBuiltinsModule("frontend", map[string]object.Object{
		"new":               object.NewBuiltin("frontend.new", frontendNew),
		"get_language":      object.NewBuiltin("frontend.get_language", frontendGetLanguage),
		"new_fontsource":    object.NewBuiltin("frontend.new_fontsource", frontendNewFontsource),
		"new_text":          object.NewBuiltin("frontend.new_text", newText),
		"new_table":         object.NewBuiltin("frontend.new_table", newTable),
		"new_tr":            object.NewBuiltin("frontend.new_tr", newTr),
		"new_td":            object.NewBuiltin("frontend.new_td", newTd),
		"font_weight_400":   object.NewInt(400),
		"font_style_normal": object.NewInt(0),
	})
}
