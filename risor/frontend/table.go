package frontend

import (
	"context"

	"github.com/boxesandglue/boxesandglue/frontend"

	rbag "github.com/boxesandglue/cli/risor/backend/bag"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const FrontendTableType = "frontend.table"
const FrontendTrType = "frontend.tr"
const FrontendTdType = "frontend.td"

type Table struct {
	Value *frontend.Table
}

type Tr struct {
	Value *frontend.TableRow
}

type Td struct {
	Value *frontend.TableCell
}

func newTable(ctx context.Context, args ...object.Object) object.Object {
	return &Table{
		Value: &frontend.Table{},
	}
}
func newTr(ctx context.Context, args ...object.Object) object.Object {
	return &Tr{
		Value: &frontend.TableRow{},
	}
}
func newTd(ctx context.Context, args ...object.Object) object.Object {
	return &Td{
		Value: &frontend.TableCell{},
	}
}

// Type of the object.
func (tbl *Table) Type() object.Type {
	return FrontendTableType
}

func (tbl *Table) append(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("frontend.table.append() takes exactly one argument")
	}
	switch args[0].Type() {
	case FrontendTrType:
		tbl.Value.Rows = append(tbl.Value.Rows, args[0].(*Tr).Value)
	default:
		return object.ArgsErrorf("frontend.table.append() expects a tr argument")
	}
	return tbl
}

// Inspect returns a string representation of the given object.
func (tbl *Table) Inspect() string {
	return "frontend.table"
}

// Interface converts the given object to a native Go value.
func (tbl *Table) Interface() interface{} {
	return nil
}

// Equals returns True if the given object is equal to this object.
func (tbl *Table) Equals(other object.Object) object.Object {
	panic("not implemented") // TODO: Implement
}

// GetAttr returns the attribute with the given name from this object.
func (tbl *Table) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "append":
		return object.NewBuiltin("append", tbl.append), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (tbl *Table) SetAttr(name string, value object.Object) error {
	switch name {
	case "max_width":
		if v, ok := value.(*rbag.RSP); ok {
			tbl.Value.MaxWidth = v.Value
			return nil
		}
	case "stretch":
		if v, ok := value.(*object.Bool); ok {
			tbl.Value.Stretch = v.IsTruthy()
			return nil
		}
	case "width":
	}
	return object.Errorf("cannot set attribute %s on table", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (tbl *Table) IsTruthy() bool {
	panic("not implemented") // TODO: Implement
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (tbl *Table) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	panic("not implemented") // TODO: Implement
}

// Cost returns the incremental processing cost of this object.
func (tbl *Table) Cost() int {
	return 0
}

// ----------------------------------------------------------
// Tr
// ----------------------------------------------------------

// Type of the object.
func (tr *Tr) Type() object.Type {
	return FrontendTrType
}

func (tr *Tr) append(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("frontend.tr.append() takes exactly one argument")
	}
	switch args[0].Type() {
	case FrontendTdType:
		tr.Value.Cells = append(tr.Value.Cells, args[0].(*Td).Value)
	default:
		return object.ArgsErrorf("frontend.tr.append() expects a td argument")
	}
	return tr
}

// Inspect returns a string representation of the given object.
func (tr *Tr) Inspect() string {
	panic("not implemented") // TODO: Implement
}

// Interface converts the given object to a native Go value.
func (tr *Tr) Interface() interface{} {
	panic("not implemented") // TODO: Implement
}

// Equals returns True if the given object is equal to this object.
func (tr *Tr) Equals(other object.Object) object.Object {
	panic("not implemented") // TODO: Implement
}

// GetAttr returns the attribute with the given name from this object.
func (tr *Tr) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "append":
		return object.NewBuiltin("append", tr.append), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (tr *Tr) SetAttr(name string, value object.Object) error {
	panic("not implemented") // TODO: Implement
}

// IsTruthy returns true if the object is considered "truthy".
func (tr *Tr) IsTruthy() bool {
	panic("not implemented") // TODO: Implement
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (tr *Tr) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	panic("not implemented") // TODO: Implement
}

// Cost returns the incremental processing cost of this object.
func (tr *Tr) Cost() int {
	return 0
}

// ----------------------------------------------------------
// Td
// ----------------------------------------------------------

// Type of the object.
func (td *Td) Type() object.Type {
	return FrontendTdType
}

func (td *Td) append(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.ArgsErrorf("frontend.td.append() takes exactly one argument")
	}
	switch args[0].Type() {
	case object.STRING:
		td.Value.Contents = append(td.Value.Contents, args[0].(*object.String).Value())
	case FrontendTextType:
		td.Value.Contents = append(td.Value.Contents, args[0].(*text).Value)
	default:
		return object.ArgsErrorf("frontend.td.append() expects a string or text argument")
	}
	return td
}

// Inspect returns a string representation of the given object.
func (td *Td) Inspect() string {
	return FrontendTdType
}

// Interface converts the given object to a native Go value.
func (td *Td) Interface() interface{} {
	return td.Value
}

// Equals returns True if the given object is equal to this object.
func (td *Td) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (td *Td) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "append":
		return object.NewBuiltin("append", td.append), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (td *Td) SetAttr(name string, value object.Object) error {
	switch name {
	case "align":
		if v, ok := value.(*object.String); ok {
			switch v.Value() {
			case "left":
				td.Value.HAlign = frontend.HAlignLeft
			case "center":
				td.Value.HAlign = frontend.HAlignCenter
			case "right":
				td.Value.HAlign = frontend.HAlignRight
			case "justify":
				td.Value.HAlign = frontend.HAlignJustified
			default:
				return object.Errorf("invalid value for align: %s", v.Value())
			}
			return nil
		}
	case "border_top_width":
		if v, ok := value.(*rbag.RSP); ok {
			td.Value.BorderTopWidth = v.Value
			return nil
		}
	case "border_bottom_width":
		if v, ok := value.(*rbag.RSP); ok {
			td.Value.BorderBottomWidth = v.Value
			return nil
		}
	case "border_left_width":
		if v, ok := value.(*rbag.RSP); ok {
			td.Value.BorderLeftWidth = v.Value
			return nil
		}
	case "border_right_width":
		if v, ok := value.(*rbag.RSP); ok {
			td.Value.BorderRightWidth = v.Value
			return nil
		}
	case "padding_top":
		if v, ok := value.(*rbag.RSP); ok {
			td.Value.PaddingTop = v.Value
			return nil
		}
	case "padding_bottom":
		if v, ok := value.(*rbag.RSP); ok {
			td.Value.PaddingBottom = v.Value
			return nil
		}
	}
	return object.Errorf("cannot set attribute %s on td", name)
}

// IsTruthy returns true if the object is considered "truthy".
func (td *Td) IsTruthy() bool {
	return true
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (td *Td) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on frontendTd", opType)
}

// Cost returns the incremental processing cost of this object.
func (td *Td) Cost() int {
	return 0
}
