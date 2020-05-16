package argmapper

import (
	"fmt"
	"reflect"
)

// Conv represents a converter function that knows how to convert
// from some set of input parameters to some set of output parameters.
//
// Converters are used if a direct match argument isn't found for a Func call.
// If a converter exists (or a chain of converts) to go from the input arguments
// to the desired argument, then the chain will be called and the result used.
//
// Converter Basics
//
// Converters must take a struct as input and return a struct as output. The
// input struct is identical to a Func and arguments are mapped directly to it.
//
// The output struct is similar to the input struct, except that the keys and
// tags of the output struct will set new values for that input type. These
// values are only set for that specific chain execution. For example:
//
//    TODO
//
// Attempted Conversions
//
// The output type can also be a pointer to a struct. If a nil pointer is
// returned, the conversion is assumed to have failed. In this case, an
// alternate chain (if it exists) will be tried.
//
//    TODO
//
// Errors
//
// A second output type of error can be used to specify any errors that
// occurred during conversion. If a non-nil error is returned, alternate
// chains will be attempted. If all chains fail, the error will be reported
// to the user. In all cases, the errors are made available in the Result type
// for logging.
type Conv struct {
	inputs  map[string]reflect.Type
	outputs map[string]reflect.Type
}

// NewConv constructs a new converter. See the docs on Conv for more info.
func NewConv(f interface{}) (*Conv, error) {
	fv := reflect.ValueOf(f)
	ft := fv.Type()
	if k := ft.Kind(); k != reflect.Func {
		return nil, fmt.Errorf("fn should be a function, got %s", k)
	}

	// We only accept zero or 1 arguments right now. In the future we
	// could potentially expand this to support multiple args that are
	// all structs we populate but for now lets just simplify this.
	if ft.NumIn() > 1 {
		return nil, fmt.Errorf("function must take one struct arg")
	}

	// Our argument must be a struct
	typ := ft.In(0)
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("function must take one struct arg")
	}

	// We expect one or two results. Either way the first result must
	// be a struct.
	if ft.NumOut() == 0 || ft.NumOut() > 2 {
		return nil, fmt.Errorf("function must return one or two results")
	}

	return nil, nil
}
