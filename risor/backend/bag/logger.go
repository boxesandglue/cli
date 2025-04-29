package bag

import (
	"context"
	"log/slog"

	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
)

const LoggerType = "slog.logger"

type logger struct {
	value *slog.Logger
}

func (l *logger) log(lvl string) func(context.Context, ...object.Object) object.Object {
	return func(ctx context.Context, args ...object.Object) object.Object {
		var fn func(string, ...any)
		switch lvl {
		case "debug":
			fn = l.value.Debug
		case "info":
			fn = l.value.Info
		case "warn":
			fn = l.value.Warn
		case "error":
			fn = l.value.Error
		default:
			return object.Errorf("unknown log level: %s", lvl)
		}

		firstArg := args[0]
		if firstArg.Type() != object.STRING {
			return object.ArgsErrorf("slog.%s() expects a string argument (message)", lvl)
		}
		var optargs []any
		if len(args) > 1 {
			for _, arg := range args[1:] {
				optargs = append(optargs, arg.Interface())
			}
		}
		fn(firstArg.(*object.String).Value(), optargs...)
		return nil
	}
}

// Type of the object.
func (l *logger) Type() object.Type {
	return LoggerType
}

// Inspect returns a string representation of the given object.
func (l *logger) Inspect() string {
	return "logger"
}

// Interface converts the given object to a native Go value.
func (l *logger) Interface() interface{} {
	return l.value
}

// Equals returns True if the given object is equal to this object.
func (l *logger) Equals(other object.Object) object.Object {
	return object.False
}

// GetAttr returns the attribute with the given name from this object.
func (l *logger) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "debug":
		return object.NewBuiltin("log.debug", l.log("debug")), true
	case "info":
		return object.NewBuiltin("log.info", l.log("info")), true
	case "warn":
		return object.NewBuiltin("log.warn", l.log("warn")), true
	case "error":
		return object.NewBuiltin("log.error", l.log("error")), true
	}
	return nil, false
}

// SetAttr sets the attribute with the given name on this object.
func (l *logger) SetAttr(name string, value object.Object) error {
	return nil
}

// IsTruthy returns true if the object is considered "truthy".
func (l *logger) IsTruthy() bool {
	return l.value != nil
}

// RunOperation runs an operation on this object with the given
// right-hand side object.
func (l *logger) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.Errorf("operation %s not supported on logger", opType)
}

// Cost returns the incremental processing cost of this object.
func (l *logger) Cost() int {
	return 0
}
