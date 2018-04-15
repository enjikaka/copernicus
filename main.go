package copernicus

// GetDemoString : Get a hello world string
func GetDemoString() string {
	coords := Coordinates{58.903839, 6.523730}
	code := coords.GetOpenLocationCode()

	return code
}
