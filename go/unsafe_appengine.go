// +build appengine

package flatbuffers

var (
	// See http://golang.org/ref/spec#Numeric_types

	// SizeUint8 is the byte size of a uint8.
	SizeUint8 int = 1
	// SizeUint16 is the byte size of a uint16.
	SizeUint16 int = 2
	// SizeUint32 is the byte size of a uint32.
	SizeUint32 int = 4
	// SizeUint64 is the byte size of a uint64.
	SizeUint64 int = 8

	// SizeInt8  is the byte size of a int8.
	SizeInt8 int = 1
	// SizeInt16 is the byte size of a int16.
	SizeInt16 int = 2
	// SizeInt32 is the byte size of a int32.
	SizeInt32 int = 4
	// SizeInt64 is the byte size of a int64.
	SizeInt64 int = 8

	// SizeFloat32 is the byte size of a float32.
	SizeFloat32 int = 4
	// SizeFloat64 is the byte size of a float64.
	SizeFloat64 int = 8

	// SizeByte is the byte size of a byte.
	// The `byte` type is aliased (by Go definition) to uint8.
	SizeByte = SizeUint8

	// SizeBool is the byte size of a bool.
	// The `bool` type is aliased (by flatbuffers convention) to uint8.
	SizeBool = SizeUint8

	// SizeSOffsetT is the byte size of an SOffsetT.
	SizeSOffsetT int = 4
	// SizeUOffsetT is the byte size of an UOffsetT.
	SizeUOffsetT int = 4
	// SizeVOffsetT is the byte size of an VOffsetT.
	SizeVOffsetT int = 2
)
