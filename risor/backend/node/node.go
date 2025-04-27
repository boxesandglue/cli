package node

import (
	"github.com/boxesandglue/boxesandglue/backend/node"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

type Node struct {
	Value node.Node
}

// Type of the object.
func (vl *Node) Type() object.Type {
	return "node.node"
}

// Inspect returns a string representation of the given object.
func (vl *Node) Inspect() string {
	return "a node"
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
	return object.Errorf("cannot set attribute %s on vlist", name)
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
