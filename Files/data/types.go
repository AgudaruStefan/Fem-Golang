package data

import "fmt"

type location string

func (origin location) DistanceTo(destination location) distance {
	fmt.Printf("Origin %v Destination %v", origin, destination)
	return 10
}

func locationTest() {
	nyc := location("32.321, 23.2123")
	DistanceTo(location("-23, -44"))
	nyc.DistanceTo(location("-23, -44"))
}

type distance float64 // miles
type distanceKm float64

func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func test() {
	d := distance(34.5)
	km := d.ToKm()
	km.ToMiles()
	print(d)
}
