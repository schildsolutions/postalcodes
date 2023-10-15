package postalcodes_test

import (
	"errors"
	"github.com/schild-media/postalcodes"
	"math"
	"testing"
)

var codeBrandenburgGate = "10117"
var codeCologneCathedral = "50667"
var distBrandenburgGateCologneCathedral float64 = 475.36
var tolerance float64 = 5

var table = postalcodes.Table{
	Codes: []postalcodes.PostalCode{
		{Code: codeBrandenburgGate, City: "Berlin", Lat: 52.517, Long: 13.3872},
		{Code: codeCologneCathedral, City: "Köln", Lat: 50.9387, Long: 6.9547},
	},
}

func init() {
	table.BuildIndex()
}

func inRange(value float64, target float64, targetRange float64) bool {
	return math.Abs(value-target) < targetRange
}

func TestGet(t *testing.T) {
	city, err := table.Get(codeBrandenburgGate)
	if err != nil {
		t.Fatalf("failed to get city: %s", err)
	}

	if city.City != "Berlin" {
		t.Fatal("failed to get city: fetched wrong city")
	}
}

func TestDistance(t *testing.T) {
	distance, err := table.CalculateDistance(codeBrandenburgGate, codeCologneCathedral)
	if err != nil {
		t.Fatalf("failed to calculate distance: %s", err)
	}

	if !inRange(distance, distBrandenburgGateCologneCathedral, tolerance) {
		t.Fatalf("distance is way too off")
	}
}

func TestNotFound(t *testing.T) {
	_, err := table.Get("not-existing")
	if !errors.Is(err, postalcodes.ErrorNotFound) {
		t.Fatal("should return ErrorNotFound")
	}
}
