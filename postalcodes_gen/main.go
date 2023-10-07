package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"
)

//go:embed template.txt
var templateFile string

const (
	POS_POSTAL_CODE = 1
	POS_CITY        = 2
	POS_LAT         = 9
	POS_LONG        = 10
)

func getFilename() string {
	if len(os.Args) == 2 {
		return os.Args[1]
	}
	return ""
}

func main() {
	filename := getFilename()
	if filename == "" {
		fmt.Println("usage: postalcodes_gen <filename>")
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read the input file: %s\n", err)
		os.Exit(1)
	}

	fileContent := string(data)
	fileLines := strings.Split(fileContent, "\n")

	var outputLines []string
	for _, line := range fileLines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "\t")
		code := parts[POS_POSTAL_CODE]
		city := parts[POS_CITY]
		lat := parts[POS_LAT]
		long := parts[POS_LONG]

		outputLines = append(outputLines, fmt.Sprintf(
			"\t\t{Code: \"%s\", City: \"%s\", Lat: %s, Long: %s},",
			code,
			city,
			lat,
			long,
		))
	}
	outputContent := strings.Join(outputLines, "\n")

	tmpl, err := template.New("template").Parse(templateFile)
	if err != nil {
		fmt.Printf("Failed to load template: %s\n", err)
		os.Exit(1)
	}

	file, err := os.Create("postalcodes.generated.go")
	if err != nil {
		fmt.Printf("Failed to open output file: %s\n", err)
		os.Exit(1)
	}

	defer func() {
		_ = file.Close()
	}()

	err = tmpl.Execute(file, outputContent)
	if err != nil {
		fmt.Printf("Failed to write output file: %s\n", err)
		os.Exit(1)
	}
}
