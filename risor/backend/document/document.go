package document

import (
	"context"
	"os"

	"github.com/boxesandglue/boxesandglue/backend/bag"
	"github.com/boxesandglue/boxesandglue/backend/document"
	"github.com/boxesandglue/boxesandglue/frontend"
	rbag "github.com/boxesandglue/cli/risor/backend/bag"
	rlang "github.com/boxesandglue/cli/risor/backend/lang"
	rnode "github.com/boxesandglue/cli/risor/backend/node"
	rpdf "github.com/boxesandglue/cli/risor/baseline-pdf"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

// Document represents a PDF document object.
type Document struct {
	PDFDoc      *document.PDFDocument
	Attachments *object.List
}

func (doc *Document) createImageNodeFromImagefile(ctx context.Context, args ...object.Object) object.Object {
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

	attachments := make([]document.Attachment, 0, doc.Attachments.Size())
	for _, entry := range doc.Attachments.Interface().([]any) {
		attachment := document.Attachment{}
		value := entry.(map[string]any)
		for k, v := range value {
			switch k {
			case "filename":
				if v, ok := v.(string); ok {
					data, err := os.ReadFile(v)
					if err != nil {
						return object.NewError(err)
					}
					attachment.Name = v
					attachment.Data = data
				}
			case "mimetype":
				if v, ok := v.(string); ok {
					attachment.MimeType = v
				}
			case "description":
				if v, ok := v.(string); ok {
					attachment.Description = v
				}
			}
		}
		bag.Logger.Info("Add attachment", "filename", attachment.Name)
		attachments = append(attachments, attachment)
	}

	doc.PDFDoc.Attachments = attachments
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

func (doc *Document) outputXMLDump(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("document.output_xml_dump() takes exactly one argument (filename)")
	}
	if args[0].Type() != object.STRING {
		return object.ArgsErrorf("document.output_xml_dump() expects a string argument (filename)")
	}
	filename := args[0].(*object.String).Value()
	if filename == "" {
		return object.ArgsErrorf("document.output_xml_dump() expects a non-empty string argument (filename)")
	}
	w, err := os.Create(filename)
	if err != nil {
		return object.NewError(err)
	}
	defer w.Close()
	if err = doc.PDFDoc.OutputXMLDump(w); err != nil {
		return object.NewError(err)
	}
	return nil
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
	case "attachments":
		return doc.Attachments, true
	case "create_image_node_from_imagefile":
		return object.NewBuiltin("document.create_image_node_from_imagefile", doc.createImageNodeFromImagefile), true
	case "filename":
		return object.NewString(doc.PDFDoc.Filename), true
	case "finish":
		return object.NewBuiltin("document.finish", doc.finish), true
	case "load_imagefile":
		return object.NewBuiltin("document.load_image_file", doc.loadImageFile), true
	case "new_page":
		return object.NewBuiltin("document.new_page", doc.newPage), true
	case "output_xml_dump":
		return object.NewBuiltin("document.output_xml_dump", doc.outputXMLDump), true
	case "pdf_writer":
		return &rpdf.PDF{Value: doc.PDFDoc.PDFWriter}, true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (doc *Document) SetAttr(name string, value object.Object) error {
	switch name {
	case "author":
		if value.Type() == object.STRING {
			doc.PDFDoc.Author = value.(*object.String).Value()
			return nil
		}
		return object.Errorf("author must be a string")
	case "bleed":
		if value.Type() == rbag.ScaledPointType {
			doc.PDFDoc.Bleed = value.(*rbag.RSP).Value
			return nil
		}
		return object.Errorf("bleed must be a bag.scaledpoint")
	case "compresslevel":
		if value.Type() == object.INT {
			doc.PDFDoc.CompressLevel = uint(value.(*object.Int).Value())
			return nil
		}
		return object.Errorf("compresslevel must be an int")
	case "creation_date":
		if value.Type() == object.TIME {
			doc.PDFDoc.CreationDate = value.(*object.Time).Value()
			return nil
		}
		return object.Errorf("creation_date must be a time")
	case "creator":
		if value.Type() == object.STRING {
			doc.PDFDoc.Creator = value.(*object.String).Value()
			return nil
		}
		return object.Errorf("creator must be a string")
	case "default_page_height":
		if value.Type() == rbag.ScaledPointType {
			doc.PDFDoc.DefaultPageHeight = value.(*rbag.RSP).Value
			return nil
		}
		return object.Errorf("default_page_height must be a bag.scaledpoint")
	case "default_page_width":
		if value.Type() == rbag.ScaledPointType {
			doc.PDFDoc.DefaultPageWidth = value.(*rbag.RSP).Value
			return nil
		}
		return object.Errorf("default_page_width must be a bag.scaledpoint")
	case "dump_output":
		if value.Type() == object.BOOL {
			doc.PDFDoc.DumpOutput = value.(*object.Bool).Value()
			return nil
		}
		return object.Errorf("dump_output must be a bool")
	case "format":
		if value.Type() == object.STRING {
			switch value.(*object.String).Value() {
			case "":
				doc.PDFDoc.Format = document.FormatPDF
			case "PDF/A-3b":
				doc.PDFDoc.Format = document.FormatPDFA3b
			case "PDF/X-3":
				doc.PDFDoc.Format = document.FormatPDFX3
			case "PDF/X-4":
				doc.PDFDoc.Format = document.FormatPDFX4
			case "PDF/UA":
				doc.PDFDoc.Format = document.FormatPDFUA
			default:
				return object.Errorf("format must be one of \"\", \"PDF/A-3b\", \"PDF/X-3\", \"PDF/X-4\", \"PDF/UA\"")
			}
			return nil
		}
		return object.Errorf("format must be a string (one of \"\", \"PDF/A-3b\", \"PDF/X-3\", \"PDF/X-4\", \"PDF/UA\")")
	case "keywords":
		if value.Type() == object.STRING {
			doc.PDFDoc.Keywords = value.(*object.String).Value()
			return nil
		}
		return object.Errorf("keywords must be a string")
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
	case "show_cutmarks":
		if value.Type() == object.BOOL {
			doc.PDFDoc.ShowCutmarks = value.(*object.Bool).Value()
			return nil
		}
		return object.Errorf("show_cutmarks must be a bool")
	case "show_hyperlinks":
		if value.Type() == object.BOOL {
			doc.PDFDoc.ShowHyperlinks = value.(*object.Bool).Value()
			return nil
		}
		return object.Errorf("show_hyperlinks must be a bool")
	case "suppressinfo":
		if value.Type() == object.BOOL {
			doc.PDFDoc.SuppressInfo = value.(*object.Bool).Value()
			return nil
		}
		return object.Errorf("suppressinfo must be a bool")
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
	case "viewer_preferences":
		if value.Type() == object.MAP {
			m := value.(*object.Map).Value()
			doc.PDFDoc.ViewerPreferences = make(map[string]string)
			for k, v := range m {
				if v.Type() != object.STRING {
					return object.Errorf("viewer_preferences must be a map of strings")
				}
				doc.PDFDoc.ViewerPreferences[k] = v.(*object.String).Value()
			}
			return nil
		}
		return object.Errorf("viewer_preferences must be a map")
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
