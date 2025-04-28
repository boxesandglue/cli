package document

import (
	"context"

	"github.com/boxesandglue/boxesandglue/backend/document"
	"github.com/boxesandglue/boxesandglue/frontend"
	"github.com/boxesandglue/cli/risor/backend/bag"
	rlang "github.com/boxesandglue/cli/risor/backend/lang"
	rnode "github.com/boxesandglue/cli/risor/backend/node"
	rpdf "github.com/boxesandglue/cli/risor/baseline-pdf"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type Document struct {
	PDFDoc *document.PDFDocument
}

func (doc *Document) CreateImageNodeFromImagefile(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.ArgsErrorf("document.create_image_node_from_imagefile() takes exactly three arguments")
	}
	firstArg := args[0]
	secondArg := args[1]
	thirdArg := args[2]

	if firstArg.Type() != "baseline-pdf.imagefile" {
		return object.ArgsErrorf("document.create_image_node_from_imagefile() expects a baseline-pdf.imagefile argument (imagefile)")
	}
	if secondArg.Type() != object.INT {
		return object.ArgsErrorf("document.create_image_node_from_imagefile() expects an int argument (page number)")
	}
	if thirdArg.Type() != object.STRING {
		return object.ArgsErrorf("document.create_image_node_from_imagefile() expects a string argument (PDF box)")
	}
	imgNode := doc.PDFDoc.CreateImageNodeFromImagefile(firstArg.(*rpdf.ImageFile).Value, int(secondArg.(*object.Int).Value()), thirdArg.(*object.String).Value())
	newnode := &rnode.Node{Value: imgNode}
	return newnode
}

func (doc *Document) newPage(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.ArgsErrorf("document.new_page() takes no arguments")
	}
	return newPage(doc.PDFDoc)
}

func (doc *Document) finish(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.ArgsErrorf("document.finish() takes no arguments")
	}
	doc.PDFDoc.Finish()
	return nil
}

func (doc *Document) loadImageFile(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("document.load_image_file() takes exactly one argument (filename)")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("document.load_image_file() expects a string argument (filename)")
	}
	filename := args[0].(*object.String).Value()
	imgf, err := doc.PDFDoc.LoadImageFile(filename)
	if err != nil {
		return object.NewError(err)
	}
	return &rpdf.ImageFile{Value: imgf}
}

// Type of the object.
func (doc *Document) Type() object.Type {
	return "backend.document"
}

// Inspect returns a string representation of the given object.
func (doc *Document) Inspect() string {
	return doc.PDFDoc.Filename
}

// Interface converts the given object to a native Go value.
func (doc *Document) Interface() interface{} {
	return doc.PDFDoc
}

// Equals returns True if the given object is equal to this object.
func (doc *Document) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (doc *Document) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "create_image_node_from_imagefile":
		return object.NewBuiltin("document.create_image_node_from_imagefile", doc.CreateImageNodeFromImagefile), true
	case "filename":
		return object.NewString(doc.PDFDoc.Filename), true
	case "finish":
		return object.NewBuiltin("document.finish", doc.finish), true
	case "load_imagefile":
		return object.NewBuiltin("document.load_image_file", doc.loadImageFile), true
	case "new_page":
		return object.NewBuiltin("document.new_page", doc.newPage), true
	}
	return nil, false
}

/*
   Attachments          []Attachment
   Bleed                bag.ScaledPoint
   ColorProfile         *ColorProfile
   CompressLevel        uint
   CreationDate         time.Time
   CurrentPage          *Page
   DefaultLanguage      *lang.Lang
   DefaultPageHeight    bag.ScaledPoint
   DefaultPageWidth     bag.ScaledPoint
   DumpOutput           bool
   Faces                []*pdf.Face
   Format               Format
   Languages            map[string]*lang.Lang
   Pages                []*Page
   RootStructureElement *StructureElement
   ShowCutmarks         bool
   ShowHyperlinks       bool
   Spotcolors           []*color.Color
   SuppressInfo         bool
   ViewerPreferences    map[string]string
*/

// SetAttr sets the attribute with the given name on this object.
func (doc *Document) SetAttr(name string, value object.Object) error {
	switch name {
	case "language":
		if value.Type() == object.STRING {
			l, err := frontend.GetLanguage(value.(*object.String).Value())
			if err != nil {
				return err
			}
			doc.PDFDoc.SetDefaultLanguage(l)
			return nil
		}
		if value.Type() == "backend.lang" {
			l := value.(*rlang.Lang).Value
			doc.PDFDoc.SetDefaultLanguage(l)
			return nil
		}
	case "author":
		if value.Type() == object.STRING {
			doc.PDFDoc.Author = value.(*object.String).Value()
			return nil
		}
		return object.Errorf("author must be a string")
	case "creator":
		if value.Type() == object.STRING {
			doc.PDFDoc.Creator = value.(*object.String).Value()
			return nil
		}
		return object.Errorf("creator must be a string")
	case "default_page_height":
		if value.Type() == bag.ScaledPointType {
			doc.PDFDoc.DefaultPageHeight = value.(*bag.RSP).Value
			return nil
		}
		return object.Errorf("default_page_height must be a bag.scaledpoint")
	case "default_page_width":
		if value.Type() == bag.ScaledPointType {
			doc.PDFDoc.DefaultPageWidth = value.(*bag.RSP).Value
			return nil
		}
		return object.Errorf("default_page_width must be a bag.scaledpoint")
	case "keywords":
		if value.Type() == object.STRING {
			doc.PDFDoc.Keywords = value.(*object.String).Value()
			return nil
		}
		return object.Errorf("keywords must be a string")
	case "subject":
		if value.Type() == object.STRING {
			doc.PDFDoc.Subject = value.(*object.String).Value()
			return nil
		}
		return object.Errorf("subject must be a string")
	case "title":
		if value.Type() == object.STRING {
			doc.PDFDoc.Title = value.(*object.String).Value()
			return nil
		}
		return object.Errorf("title must be a string")
	}
	return object.Errorf("cannot set attribute %s on document", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (doc *Document) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (doc *Document) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on document", opType)
}

// Cost returns the incremental processing cost of this object.
func (doc *Document) Cost() int {
	return 0
}
