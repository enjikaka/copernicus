package copernicus

import "testing"

func TestDemoString(t *testing.T) {
	v := GetDemoString()

	if v != "9FC8WG3F+GF" {
		t.Error("Expected '9FC8WG3F+GF', got ", v)
	}
}
