package node

import (
	"context"

	"github.com/boxesandglue/boxesandglue/backend/node"
	rbag "github.com/boxesandglue/cli/risor/backend/bag"
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

// Type of the object.
func (vl *Node) Type() object.Type {
	switch vl.Value.Type() {
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
func (vl *Node) Inspect() string {
	return node.String(vl.Value)
}

// Interface converts the given object to a native Go value.
func (vl *Node) Interface() interface{} {
	return vl.Value
}

// Equals returns True if the given object is equal to this object.
func (vl *Node) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (vl *Node) GetAttr(name string) (object.Object, bool) {
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (vl *Node) SetAttr(name string, value object.Object) error {
	switch name {
	case "width":
		if value.Type() != "bag.scaledpoint" {
			return object.ArgsErrorf("node.width() expects a bag.scaledpoint argument")
		}
		val := value.(*rbag.RSP)

		switch t := vl.Value.(type) {
		case *node.Image:
			t.Width = val.Value
			return nil
		}
	case "height":
		if value.Type() != "bag.scaledpoint" {
			return object.ArgsErrorf("node.height() expects a bag.scaledpoint argument")
		}
		val := value.(*rbag.RSP)
		switch t := vl.Value.(type) {
		case *node.Image:
			t.Height = val.Value
			return nil
		}
	}
	return object.Errorf("cannot set attribute %s on node", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (vl *Node) IsTruthy() bool {
	return vl.Value != nil
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (vl *Node) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on vlist", opType)
}

// Cost returns the incremental processing cost of this object.
func (vl *Node) Cost() int {
	return 0
}

// Module returns the node module.
func Module() *object.Module {
	return object.NewBuiltinsModule("node", map[string]object.Object{
		"vpack": object.NewBuiltin("node.vpack", vpack),
		"debug": object.NewBuiltin("node.debug", debug),
	})
}
