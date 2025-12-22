package helper

import "bytes"

func ParseToString(byteArray []byte, length int) string {
	n := bytes.Index(byteArray, []byte{0})
	if n == -1 || n > length-1 {
		n = length - 1
	}
	return string(byteArray[:n])
}

func BoolToUint16(value bool) uint16 {
	if value {
		return 1
	}
	return 0
}

func Uint16ToBool(value uint16) bool {
	return value == 1
}
