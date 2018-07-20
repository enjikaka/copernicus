package copernicus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

type IdentifyResult struct {
	Results []struct {
		LayerID          int    `json:"layerId"`
		LayerName        string `json:"layerName"`
		DisplayFieldName string `json:"displayFieldName"`
		Value            string `json:"value"`
		Attributes       struct {
			OBJECTID    string `json:"OBJECTID"`
			Shape       string `json:"Shape"`
			Code12      string `json:"code_12"`
			ID          string `json:"ID"`
			Remark      string `json:"Remark"`
			AreaHa      string `json:"Area_Ha"`
			ShapeLength string `json:"Shape_Length"`
			ShapeArea   string `json:"Shape_Area"`
		} `json:"attributes"`
	} `json:"results"`
}

type Fetcher struct{}

type IdentifySearchQuery struct {
	Geometry       string    `url:"geometry"`
	GeometryType   string    `url:"geometryType"`
	Tolerance      int       `url:"tolerance"`
	MapExtent      []float64 `url:"mapExtent"`
	ReturnGeometry bool      `url:"returnGeometry"`
	ImageDisplay   []int     `url:"imageDisplay"`
	Format         string    `url:"f"`
}

// BuildSearchQuery : Builds search query for the Identify endpoint on the copernicus MapServer
func BuildSearchQuery(coords Coordinates) string {
	geometry := coords.GetBoundsInMeters()

	json, err := json.Marshal(geometry)
	if err != nil {
		log.Fatal(err)
	}
	geometryString := string(json)

	identifySearchQuery := IdentifySearchQuery{
		Geometry:       geometryString,
		GeometryType:   "esriGeometryEnvelope",
		Tolerance:      1,
		MapExtent:      []float64{geometry.XMin, geometry.YMin, geometry.XMax, geometry.YMax},
		ReturnGeometry: false,
		ImageDisplay:   []int{10, 10},
		Format:         "pjson",
	}

	v, _ := query.Values(identifySearchQuery)
	return v.Encode()
}

// Identify : Fetches the land cover information for the coordinate
func (f Fetcher) Identify(coords Coordinates) (IdentifyResult, error) {
	searchQuery := BuildSearchQuery(coords)
	url := fmt.Sprintf("https://copernicus.discomap.eea.europa.eu/arcgis/rest/services/Corine/CLC2012_WM/MapServer/identify?%s", searchQuery)

	spaceClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	result := IdentifyResult{}
	jsonErr := json.Unmarshal(body, &result)

	return result, jsonErr
}
