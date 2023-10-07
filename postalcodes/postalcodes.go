package postalcodes

import (
	"errors"
	"math"
)

var (
	// ErrorNotFound occurs when a postal code is requested, but could not be found
	ErrorNotFound = errors.New("postal code not found")
)

// PostalCode is one entry in the table. It contains the name of the city and the gps position.
type PostalCode struct {
	// Code is the postal code of a city's postal region
	Code string
	// City contains the name of the city
	City string
	// Lat contains the coordinate latitude
	Lat float64
	// Long contains the coordinates longitude
	Long float64
}

// Table contains all (generated) postal code entries and contains helping functions to manage those data.
type Table struct {
	Codes []PostalCode
	index map[string]*PostalCode
}

// BuildIndex builds a search index for the table.
// It needs to be called once, which is usually done in the generated file.
func (t *Table) BuildIndex() {
	t.index = make(map[string]*PostalCode)

	for _, code := range t.Codes {
		codeCopy := code
		t.index[code.Code] = &codeCopy
	}
}

// Get searches an entry by the postal code.
// If the postal code could not be found, ErrorNotFound will be returned as error.
func (t *Table) Get(postalCode string) (PostalCode, error) {
	if postalCode, ok := t.index[postalCode]; ok {
		return *postalCode, nil
	}

	return PostalCode{}, ErrorNotFound
}

// CalculateDistance takes two postal codes, searches for the gps positions and calculates the distance in kilometers.
// It uses the "Equirectangular approximation" described in this article:
// https://www.movable-type.co.uk/scripts/latlong.html
func (t *Table) CalculateDistance(postalCode1 string, postalCode2 string) (float64, error) {
	p1, err := t.Get(postalCode1)
	if err != nil {
		return 0, err
	}
	p2, err := t.Get(postalCode2)
	if err != nil {
		return 0, err
	}

	_ = p1
	_ = p2

	earthRadius := 6371.0
	lat1 := p1.Lat * math.Pi / 180
	lat2 := p2.Lat * math.Pi / 180
	long1 := p1.Long * math.Pi / 180
	long2 := p2.Long * math.Pi / 180
	x := (long2 - long1) * math.Cos((lat1+lat2)/2)
	y := lat2 - lat1
	d := math.Sqrt(x*x+y*y) * earthRadius

	return d, nil
}
