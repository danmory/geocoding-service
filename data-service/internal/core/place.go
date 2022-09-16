package core

type PlaceCoordinates struct {
	Lat  float32 `json:"lat"`
	Lon  float32 `json:"lon"`
}

type PlaceName struct {
	Name string  `json:"name"`
}

type Place struct {
	PlaceName
	PlaceCoordinates
}
