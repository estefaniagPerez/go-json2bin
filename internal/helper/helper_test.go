package helper

import "testing"

func TestBoolToUint16(t *testing.T) {
	if BoolToUint16(true) != 1 {
		t.Error("Expected 1, got 0")
	}
	if BoolToUint16(false) != 0 {
		t.Error("Expected 0, got 1")
	}
}

func TestUint16ToBool(t *testing.T) {
	if !Uint16ToBool(1) {
		t.Error("Expected true, got false")
	}
	if Uint16ToBool(0) {
		t.Error("Expected false, got true")
	}
}

func TestParseToString(t *testing.T) {
	byteArray := []byte{'h', 'e', 'l', 'l', 'o', 0, 'w', 'o', 'r', 'l', 'd'}
	result := ParseToString(byteArray, 11)
	expected := "hello"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	byteArrayNoNull := []byte{'h', 'e', 'l', 'l', 'o'}
	result = ParseToString(byteArrayNoNull, 6)
	expected = "hello"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	byteArrayFull := []byte{'h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd'}
	result = ParseToString(byteArrayFull, 11)
	expected = "helloworld"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
