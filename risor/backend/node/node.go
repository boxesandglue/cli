package node

import (
	"context"

	"github.com/boxesandglue/boxesandglue/backend/node"
	rbag "github.com/boxesandglue/cli/risor/backend/bag"
	rlang "github.com/boxesandglue/cli/risor/backend/lang"
	pdf "github.com/boxesandglue/cli/risor/baseline-pdf"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

// Node represents a node object in the risor language.
// It is a wrapper around the backend node object.
type Node struct {
	Value node.Node
}

func debug(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("node.debug() takes exactly one argument")
	}
	firstArg := args[0]
	var n *Node
	var ok bool
	if n, ok = firstArg.(*Node); !ok {
		return object.ArgsErrorf("node.debug() expects a node.node argument")
	}
	if n.Value == nil {
		return object.ArgsErrorf("node.debug() expects a non-nil node.node argument")
	}
	node.Debug(n.Interface().(node.Node))
	return nil
}

func vpack(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("node.vpack() takes exactly one argument")
	}
	firstArg := args[0]
	var n *Node
	var ok bool

	if n, ok = firstArg.(*Node); !ok {
		return object.ArgsErrorf("node.vpack() expects a node.node argument")
	}
	if n.Value == nil {
		return object.ArgsErrorf("node.vpack() expects a non-nil node.node argument")
	}
	return &Node{Value: node.Vpack(n.Interface().(node.Node))}
}

func newNode(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("node.new() takes exactly one argument")
	}
	firstArg := args[0]
	var name string

	if firstArg.Type() == object.STRING {
		name = firstArg.(*object.String).Value()
		switch name {
		case "disc":
			return &Node{Value: node.NewDisc()}
		case "glue":
			return &Node{Value: node.NewGlue()}
		case "glyph":
			return &Node{Value: node.NewGlyph()}
		case "hlist":
			return &Node{Value: node.NewHList()}
		case "image":
			return &Node{Value: node.NewImage()}
		case "kern":
			return &Node{Value: node.NewKern()}
		case "lang":
			return &Node{Value: node.NewLang()}
		case "penalty":
			return &Node{Value: node.NewPenalty()}
		case "rule":
			return &Node{Value: node.NewRule()}
		case "startstop":
			return &Node{Value: node.NewStartStop()}
		case "vlist":
			return &Node{Value: node.NewVList()}
		default:
			return object.ArgsErrorf("node.new() expects a string argument (node type), one of disc, glue, glyph, hlist, image, kern, lang, penalty, rule, startstop or vlist")
		}
	}
	return object.ArgsErrorf("node.new() expects a string argument (node type)")
}

func insertAfter(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.ArgsErrorf("node.insert_after() takes exactly three arguments")
	}
	firstArg := args[0]
	secondArg := args[1]
	thirdArg := args[2]
	var head, cur, insert *Node
	var ok bool
	if head, ok = firstArg.(*Node); !ok {
		return object.ArgsErrorf("node.insert_after() expects a node.node argument (head)")
	}
	if cur, ok = secondArg.(*Node); !ok {
		return object.ArgsErrorf("node.insert_after() expects a node.node argument (cur)")
	}
	if insert, ok = thirdArg.(*Node); !ok {
		return object.ArgsErrorf("node.insert_after() expects a node.node argument (insert)")
	}
	if head.Value == nil {
		return object.ArgsErrorf("node.insert_after() expects a non-nil node.node argument (head)")
	}
	if cur.Value == nil {
		return object.ArgsErrorf("node.insert_after() expects a non-nil node.node argument (cur)")
	}
	if insert.Value == nil {
		return object.ArgsErrorf("node.insert_after() expects a non-nil node.node argument (insert)")
	}
	n := node.InsertAfter(head.Value, cur.Value, insert.Value)
	return &Node{Value: n}
}

func insertBefore(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.ArgsErrorf("node.insert_before() takes exactly three arguments")
	}
	firstArg := args[0]
	secondArg := args[1]
	thirdArg := args[2]
	var head, cur, insert *Node
	var ok bool
	if head, ok = firstArg.(*Node); !ok {
		return object.ArgsErrorf("node.insert_before() expects a node.node argument (head)")
	}
	if cur, ok = secondArg.(*Node); !ok {
		return object.ArgsErrorf("node.insert_before() expects a node.node argument (cur)")
	}
	if insert, ok = thirdArg.(*Node); !ok {
		return object.ArgsErrorf("node.insert_before() expects a node.node argument (insert)")
	}
	if head.Value == nil {
		return object.ArgsErrorf("node.insert_before() expects a non-nil node.node argument (head)")
	}
	if cur.Value == nil {
		return object.ArgsErrorf("node.insert_before() expects a non-nil node.node argument (cur)")
	}
	if insert.Value == nil {
		return object.ArgsErrorf("node.insert_before() expects a non-nil node.node argument (insert)")
	}
	n := node.InsertBefore(head.Value, cur.Value, insert.Value)
	return &Node{Value: n}
}

func tail(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("node.tail() takes exactly one argument")
	}
	firstArg := args[0]
	var n *Node
	var ok bool
	if n, ok = firstArg.(*Node); !ok {
		return object.ArgsErrorf("node.tail() expects a node.node argument")
	}
	if n.Value == nil {
		return object.ArgsErrorf("node.tail() expects a non-nil node.node argument")
	}
	return &Node{Value: node.Tail(n.Value)}
}
func copyList(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("node.copy_list() takes exactly one argument")
	}
	firstArg := args[0]
	var n *Node
	var ok bool
	if n, ok = firstArg.(*Node); !ok {
		return object.ArgsErrorf("node.copy_list() expects a node.node argument")
	}
	if n.Value == nil {
		return object.ArgsErrorf("node.copy_list() expects a non-nil node.node argument")
	}
	return &Node{Value: node.CopyList(n.Value)}
}

// Type of the object.
func (n *Node) Type() object.Type {
	switch n.Value.Type() {
	case node.TypeDisc:
		return "node.disc"
	case node.TypeGlue:
		return "node.glue"
	case node.TypeGlyph:
		return "node.glyph"
	case node.TypeHList:
		return "node.hlist"
	case node.TypeImage:
		return "node.image"
	case node.TypeKern:
		return "node.kern"
	case node.TypeLang:
		return "node.lang"
	case node.TypePenalty:
		return "node.penalty"
	case node.TypeRule:
		return "node.rule"
	case node.TypeStartStop:
		return "node.startstop"
	case node.TypeVList:
		return "node.vlist"
	default:
		return "node.node"
	}
}

// Inspect returns a string representation of the given object.
func (n *Node) Inspect() string {
	return node.String(n.Value)
}

// Interface converts the given object to a native Go value.
func (n *Node) Interface() interface{} {
	return n.Value
}

// Equals returns True if the given object is equal to this object.
func (n *Node) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (n *Node) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "next":
		if n.Value.Next() == nil {
			return object.Nil, true
		}
		return &Node{Value: n.Value.Next()}, true
	case "prev":
		if n.Value.Prev() == nil {
			return object.Nil, true
		}
		return &Node{Value: n.Value.Prev()}, true
	}
	return nil, false
}

/*

Shift     bag.ScaledPoint // The displacement perpendicular to the progressing direction. Not used.
ShiftX   bag.ScaledPoint
ShipoutCallback StartStopFunc
Shrink       bag.ScaledPoint // The shrinkability of the glue, where width minus shrink = minimum width.
ShrinkOrder  GlueOrder
StartNode       *StartStop
Stretch      bag.ScaledPoint // The stretchability of the glue, where width plus stretch = maximum width.
StretchOrder GlueOrder       // The order of infinity of stretching.
Subtype      GlueSubtype

VAlign    VerticalAlignment
Value any
Width bag.ScaledPoint
YOffset bag.ScaledPoint
*/
// SetAttr sets the attribute with the given name on this object.
func (n *Node) SetAttr(name string, value object.Object) error {
	switch name {
	case "next":
		var otherNode *Node
		var ok bool
		if otherNode, ok = value.(*Node); !ok {
			return object.ArgsErrorf("node.next expects a node.node value")
		}
		n.Value.SetNext(otherNode.Value)
		return nil
	case "prev":
		var otherNode *Node
		var ok bool
		if otherNode, ok = value.(*Node); !ok {
			return object.ArgsErrorf("node.prev expects a node.node value")
		}
		n.Value.SetPrev(otherNode.Value)
		return nil
	case "action":
		return object.ArgsErrorf("node.action is not implemented")
	case "badness":
		if value.Type() != object.INT {
			return object.ArgsErrorf("node.badness expects an int value")
		}
		val := value.(*object.Int)
		switch t := n.Value.(type) {
		case *node.HList:
			t.Badness = int(val.Value())
			return nil
		}
	case "codepoint":
		if value.Type() != object.INT {
			return object.ArgsErrorf("node.codepoint() expects an int argument")
		}
		val := value.(*object.Int)
		switch t := n.Value.(type) {
		case *node.Glyph:
			t.Codepoint = int(val.Value())
		}
	case "components":
		if value.Type() != object.STRING {
			return object.ArgsErrorf("node.components() expects a string argument")
		}
		val := value.(*object.String)
		switch t := n.Value.(type) {
		case *node.Glyph:
			t.Components = val.Value()
		}
	case "depth":
		if value.Type() != "bag.scaledpoint" {
			return object.ArgsErrorf("node.depth() expects a bag.scaledpoint argument")
		}
		val := value.(*rbag.RSP)
		switch t := n.Value.(type) {
		case *node.Glyph:
			t.Depth = val.Value
			return nil
		case *node.HList:
			t.Depth = val.Value
			return nil
		case *node.Rule:
			t.Depth = val.Value
			return nil
		case *node.VList:
			t.Depth = val.Value
			return nil
		}
	case "font":
		// not implemented
		return object.ArgsErrorf("node.font() is not implemented")
	case "glue_order":
		if value.Type() != object.INT {
			return object.ArgsErrorf("node.glue_order expects an int value")
		}
		val := value.(*object.Int)
		switch t := n.Value.(type) {
		case *node.HList:
			t.GlueOrder = node.GlueOrder(val.Value())
			return nil
		}
	case "glue_set":
		if value.Type() != object.FLOAT {
			return object.ArgsErrorf("node.glue_set expects a float value")
		}
		val := value.(*object.Float)
		switch t := n.Value.(type) {
		case *node.HList:
			t.GlueSet = val.Value()
			return nil
		}
	case "glue_sign":
		if value.Type() != object.INT {
			return object.ArgsErrorf("node.glue_sign expects an int value")
		}
		val := value.(*object.Int)
		switch t := n.Value.(type) {
		case *node.HList:
			t.GlueSign = uint8(val.Value())
			return nil
		}
	case "height":
		if value.Type() != "bag.scaledpoint" {
			return object.ArgsErrorf("node.height expects a bag.scaledpoint value")
		}
		val := value.(*rbag.RSP)
		switch t := n.Value.(type) {
		case *node.Image:
			t.Height = val.Value
			return nil
		case *node.Glyph:
			t.Height = val.Value
			return nil
		case *node.HList:
			t.Height = val.Value
			return nil
		case *node.Rule:
			t.Height = val.Value
			return nil
		case *node.VList:
			t.Height = val.Value
			return nil
		}
	case "hide":
		if value.Type() != object.BOOL {
			return object.ArgsErrorf("node.hide expects a bool value")
		}
		val := value.(*object.Bool)
		switch t := n.Value.(type) {
		case *node.Rule:
			t.Hide = val.Value()
			return nil
		}
	case "hyphenate":
		if value.Type() != object.BOOL {
			return object.ArgsErrorf("node.hyphenate expects a bool value")
		}
		val := value.(*object.Bool)
		switch t := n.Value.(type) {
		case *node.Glyph:
			t.Hyphenate = val.Value()
			return nil
		}
	case "imagefile":
		if value.Type() != "baseline-pdf.imagefile" {
			return object.ArgsErrorf("node.imagefile expects a pdf.imagefile value")
		}
		val := value.(*pdf.ImageFile)
		switch t := n.Value.(type) {
		case *node.Image:
			t.ImageFile = val.Value
			return nil
		}
	case "kern":
		if value.Type() != "bag.scaledpoint" {
			return object.ArgsErrorf("node.kern expects a bag.scaledpoint value")
		}
		val := value.(*rbag.RSP)
		switch t := n.Value.(type) {
		case *node.Kern:
			t.Kern = val.Value
			return nil
		}
	case "lang":
		if value.Type() != "backend.lang" {
			return object.ArgsErrorf("node.lang expects a backend.lang value")
		}
		val := value.(*rlang.Lang)
		switch t := n.Value.(type) {
		case *node.Lang:
			t.Lang = val.Value
			return nil
		}
	case "list":
		if value.Type() != "node.hlist" && value.Type() != "node.vlist" {
			return object.ArgsErrorf("node.list expects a node.hlist or node.vlist value")
		}
		val := value.(*Node)
		switch t := n.Value.(type) {
		case *node.HList:
			t.List = val.Value
			return nil
		case *node.VList:
			t.List = val.Value
			return nil
		}
	case "page_number":
		if value.Type() != object.INT {
			return object.ArgsErrorf("node.page_number expects an int value")
		}
		val := value.(*object.Int)
		switch t := n.Value.(type) {
		case *node.Image:
			t.PageNumber = int(val.Value())
			return nil
		}
	case "penalty":
		if value.Type() != object.INT {
			return object.ArgsErrorf("node.penalty expects an int value")
		}
		val := value.(*object.Int)
		switch t := n.Value.(type) {
		case *node.Penalty:
			t.Penalty = int(val.Value())
			return nil
		}
	case "position":
		// not implemented
		return object.ArgsErrorf("node.position is not implemented")
	case "shift_x":
		if value.Type() != "bag.scaledpoint" {
			return object.ArgsErrorf("node.shift_x expects a bag.scaledpoint value")
		}
		val := value.(*rbag.RSP)
		switch t := n.Value.(type) {
		case *node.VList:
			t.ShiftX = val.Value
			return nil
		}
	case "used":
		if value.Type() != object.BOOL {
			return object.ArgsErrorf("node.used expects a bool value")
		}
		val := value.(*object.Bool)
		switch t := n.Value.(type) {
		case *node.Image:
			t.Used = val.Value()
			return nil
		}
	case "width":
		if value.Type() != "bag.scaledpoint" {
			return object.ArgsErrorf("node.width expects a bag.scaledpoint value")
		}
		val := value.(*rbag.RSP)

		switch t := n.Value.(type) {
		case *node.Image:
			t.Width = val.Value
			return nil
		}

	}
	return object.Errorf("cannot set attribute %s on node", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (n *Node) IsTruthy() bool {
	return n.Value != nil
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (n *Node) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on vlist", opType)
}

// Cost returns the incremental processing cost of this object.
func (n *Node) Cost() int {
	return 0
}

// Module returns the node module.
func Module() *object.Module {
	return object.NewBuiltinsModule("node", map[string]object.Object{
		"debug":         object.NewBuiltin("node.debug", debug),
		"new":           object.NewBuiltin("node.new", newNode),
		"vpack":         object.NewBuiltin("node.vpack", vpack),
		"insert_after":  object.NewBuiltin("node.insert_after", insertAfter),
		"insert_before": object.NewBuiltin("node.insert_before", insertBefore),
		"tail":          object.NewBuiltin("node.tail", tail),
		"copy_list":     object.NewBuiltin("node.copy_list", copyList),
	})
}
