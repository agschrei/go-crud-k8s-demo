package models

type Aircraft struct {
	Code  string
	Model string
	Range string
}

type LatLonCoordinate struct {
	Latitude  float64
	Longitude float64
}

type Airport struct {
	Code        string
	Name        string
	City        string
	Coordinates *LatLonCoordinate
	Timezone    string
}
