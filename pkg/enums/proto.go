package enums

// Enum : returns the protobuf enum object given its corresponding string. If the string is not found, returns the
// given default enum.
// Usage: fooEnum := goUtils.Enum("foo", pb.SomeEnum_value, pb.SomeEnum_DEFAULT)
// `pb.SomeEnum_value` is the proto-generated map
func Enum[T ~string, PB ~int32](key T, pbMap map[string]int32, defaultEnum PB) PB {
	value, ok := pbMap[string(key)]
	if !ok {
		return defaultEnum
	}
	return PB(value)
}
