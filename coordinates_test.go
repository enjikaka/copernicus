package copernicus

import "testing"

func TestGetOpenLocationCode(t *testing.T) {
	c := Coordinates{58.903839, 6.523730}
	v := c.GetOpenLocationCode()

	if v != "9FC8WG3F+GF" {
		t.Error("Expected 'hello world copernicus', got ", v)
	}
}

func TestGetBoundsInMeters(t *testing.T) {
	c := Coordinates{58.903839, 6.523730}
	xmin, ymin, xmax, ymax := c.GetBoundsInMeters()

	if xmin != 726206.6130251817 {
		t.Error("Expected '726206.6130251817', got ", xmin)
	}

	if ymin != 8.159612583541271e+06 {
		t.Error("Expected 'ymin', got ", 8.159612583541271e+06)
	}

	if xmax != 726220.5279615285 {
		t.Error("Expected 'xmax', got ", 726220.5279615285)
	}

	if ymax != 8.159639525600318e+06 {
		t.Error("Expected 'ymax', got ", 8.159639525600318e+06)
	}
}
