package postalcodes

import (
	"errors"
	"math"
)

var (
	ErrorNotFound = errors.New("postal code not found")
)

type PostalCode struct {
	Code string
	City string
	Lat  float64
	Long float64
}

type Table struct {
	Codes []PostalCode
	index map[string]*PostalCode
}

func (t *Table) BuildIndex() {
	t.index = make(map[string]*PostalCode)

	for _, code := range t.Codes {
		codeCopy := code
		t.index[code.Code] = &codeCopy
	}
}

func (t *Table) Get(postalCode string) (PostalCode, error) {
	if postalCode, ok := t.index[postalCode]; ok {
		return *postalCode, nil
	}

	return PostalCode{}, ErrorNotFound
}

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

	// more info, see: https://www.movable-type.co.uk/scripts/latlong.html
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
