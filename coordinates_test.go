package copernicus

import "testing"

func TestGetOpenLocationCode(t *testing.T) {
	c := Coordinates{58.903839, 6.523730}
	v := c.GetOpenLocationCode()

	if v != "9FC8WG3F+GF" {
		t.Error("Expected 'hello world copernicus', got ", v)
	}
}

func TestDegreesToMeters(t *testing.T) {
	x, y := DegreesToMeters(58.801089, 13.219594)

	if x != 6.54570728465879e+06 {
		t.Error("Expected '726206.6130251817', got ", x)
	}

	if y != 1.4848315377220341e+06 {
		t.Error("Expected '726206.6130251817', got ", y)
	}
}

func TestGetBoundsInMeters(t *testing.T) {
	c := Coordinates{58.903839, 6.523730}
	geometry := c.GetBoundsInMeters()

	if geometry.XMin != 726206.6130251817 {
		t.Error("Expected '726206.6130251817', got ", geometry.XMin)
	}

	if geometry.YMin != 8.159612583541271e+06 {
		t.Error("Expected '8.159612583541271e+06', got ", geometry.YMin)
	}

	if geometry.XMax != 726220.5279615285 {
		t.Error("Expected '726220.5279615285', got ", geometry.XMax)
	}

	if geometry.YMax != 8.159639525600318e+06 {
		t.Error("Expected '8.159639525600318e+06', got ", geometry.YMax)
	}
}
