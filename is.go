package objx

import (
	"reflect"
)

// IsNil gets whether the data is nil or not.
func (o *O) IsNil() bool {
	return o == nil || o.Obj() == nil
}

// IsKind gets whether the object contained is of the
// specified reflect.Kind or not.
func (o *O) IsKind(kind reflect.Kind) bool {
	return o.Kind() == kind
}

// IsMap gets whether the object contained is a map
// or not.
func (o *O) IsMap() bool {
	return o.IsKind(reflect.Map)
}

// IsList gets whether the object contained is a slice
// or array.
//
// If you care about the difference, you should use the
// IsSlice and IsArray methods instead.
func (o *O) IsList() bool {
	return o.Kind() == reflect.Slice || o.Kind() == reflect.Array
}

// IsSlice gets whether the object contained is a slice
// or not.  It is recommended that you use IsList instead.
func (o *O) IsSlice() bool {
	return o.IsKind(reflect.Slice)
}

// IsArray gets whether the object contained is a slice
// or not.  It is recommended that you use IsList instead.
func (o *O) IsArray() bool {
	return o.IsKind(reflect.Array)
}

// Gets whether the object contained is of this type
// or not.
func (o *O) IsFunc() bool {
	return o.IsKind(reflect.Func)
}
