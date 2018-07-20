package copernicus

import (
	"log"
	"math"

	"github.com/google/open-location-code/go"
)

// Coordinates : Latitude and longitude pair
type Coordinates struct {
	Lat, Lng float64
}

type Geometry struct {
	XMin             float64 `url:"xmin"`
	YMin             float64 `url:"ymin"`
	XMax             float64 `url:"xmax"`
	YMax             float64 `url:"ymax"`
	SpatialReference int     `url:"spatialReference"`
}

// DegreesToMeters : Converts the lat/lng degree pair to WSG84 meters
func DegreesToMeters(longitude float64, latitude float64) (float64, float64) {
	x := longitude * 20037508.34 / 180
	y := math.Log(math.Tan((90+latitude)*math.Pi/360)) / (math.Pi / 180)

	y = y * 20037508.34 / 180

	return x, y
}

// GetOpenLocationCode : Get the Open Location Code the coordinate is in
func (coords Coordinates) GetOpenLocationCode() string {
	return olc.Encode(coords.Lat, coords.Lng, 0)
}

// GetBoundsInMeters : Get the bounding box in meters of the open location code the coordinates are in
func (coords Coordinates) GetBoundsInMeters() Geometry {
	code := coords.GetOpenLocationCode()
	codeArea, err := olc.Decode(code)

	if err != nil {
		log.Fatal(err)
	}

	xmin, ymin := DegreesToMeters(codeArea.LngLo, codeArea.LatLo)
	xmax, ymax := DegreesToMeters(codeArea.LngHi, codeArea.LatHi)

	return Geometry{XMin: xmin, YMin: ymin, XMax: xmax, YMax: ymax, SpatialReference: 102100}
}
