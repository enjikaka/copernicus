package copernicus

import (
	"testing"
)

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
