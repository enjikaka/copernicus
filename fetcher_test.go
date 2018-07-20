package copernicus

import (
	"testing"
)

func TestsearchQuery(t *testing.T) {
	c := Coordinates{58.903839, 6.523730}
	q := BuildSearchQuery(c)

	if q != "derp" {
		t.Error("Expected 'derp', got ", q)
	}
}

func TestIdentify(t *testing.T) {
	f := Fetcher{}
	c := Coordinates{58.903839, 6.523730}
	identification, err := f.Identify(c)

	if err != nil {
		t.Error(err)
	}

	numberOfResults := len(identification.Results)

	if numberOfResults == 0 {
		t.Error("Expected 'Results' to have length larget than 0, got ", numberOfResults)
	}
}
