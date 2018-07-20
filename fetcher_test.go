package copernicus

import (
	"fmt"
	"testing"
)

func TestSearchQuery(t *testing.T) {
	c := Coordinates{58.903839, 6.523730}
	q := BuildSearchQuery(c)

	expected := "f=pjson&geometry=%7B%22xmin%22%3A726206.6130251817%2C%22ymin%22%3A8159612.583541271%2C%22xmax%22%3A726220.5279615285%2C%22ymax%22%3A8159639.525600318%2C%22spatialReference%22%3A102100%7D&geometryType=esriGeometryEnvelope&imageDisplay=10&imageDisplay=10&mapExtent=726206.6130251817&mapExtent=8.159612583541271e%2B06&mapExtent=726220.5279615285&mapExtent=8.159639525600318e%2B06&returnGeometry=false&tolerance=1"

	if q != expected {
		t.Error(fmt.Sprintf("Expected '%s', got: %s", expected, q))
	}
}

func TestIdentify(t *testing.T) {
	f := Fetcher{}
	c := Coordinates{58.903839, 6.523730}
	identification, err := f.Identify(c)

	t.Log(identification)

	if err != nil {
		t.Error(err)
	}

	numberOfResults := len(identification.Results)

	if numberOfResults == 0 {
		t.Error("Expected 'Results' to have length larget than 0, got ", numberOfResults)
	}
}
