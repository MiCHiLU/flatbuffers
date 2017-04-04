// +build appengine

package flatbuffers

// byteSliceToString same as string built-in function.
func byteSliceToString(b []byte) string {
	return string(b)
}
