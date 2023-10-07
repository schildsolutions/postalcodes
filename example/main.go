package main

import (
	"log"
)

func main() {
	const firstPostalCode = "10117"
	const secondPostalCode = "50667"

	firstCity, err := PostalCodes.Get(firstPostalCode)
	if err != nil {
		log.Fatalf("An error occurred: %s\n", err)
	}

	secondCity, err := PostalCodes.Get(secondPostalCode)
	if err != nil {
		log.Fatalf("An error occurred: %s\n", err)
	}

	distance, err := PostalCodes.CalculateDistance(firstPostalCode, secondPostalCode)
	if err != nil {
		log.Fatalf("An error occurred: %s\n", err)
	}
	log.Printf("The distance between %s and %s is %.2fkm\n", firstCity.City, secondCity.City, distance)
}
