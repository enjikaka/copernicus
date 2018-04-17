package copernicus

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

// Identify : Fetches the land cover information for the coordinate
func (f Fetcher) Identify(coords Coordinates) (IdentifyResult, error) {
	url := "https://copernicus.discomap.eea.europa.eu/arcgis/rest/services/Corine/CLC2012_WM/MapServer/identify?geometry=%7B%22xmin%22%3A1197088.0590151804%2C%22ymin%22%3A8381365.272022153%2C%22xmax%22%3A1197101.9739515274%2C%22ymax%22%3A8381393.032636984%2C%22spatialReference%22%3A102100%7D&geometryType=esriGeometryEnvelope&tolerance=1&mapExtent=1197088.0590151804%2C8381365.272022153%2C1197101.9739515274%2C8381393.032636984&returnGeometry=false&imageDisplay=10%2C10&f=pjson"

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
