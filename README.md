# go-german-postalcodes

This Go library is build to make working with postal codes simply.

## Where do I get the data from?

This library is optimized for using the german postal code
information from [GeoNames](#).

You need to download those data from:
[https://download.geonames.org/export/zip/](https://download.geonames.org/export/zip/)
(The DE.zip file).

Those data are under [CC BY 4.0 DEED](https://creativecommons.org/licenses/by/4.0/) license.
Do not forget to attribute GeoNames accordingly in your application
when you use those data.

## Generate the table

Go to your work directory where a `postalcodes.generated.go` should be generated.
Execute the postalcodes_gen tool and select the `txt` file from GeoNames.

    go run github.com/schild-media/go-geerman-postalcodes/postalcodes_gen <filename>

Example:

    go run github.com/schild-media/go-geerman-postalcodes/postalcodes_gen DE.txt

This will generate the file `postalcodes.generated.go`, which contains an instance called `PostalCodes`.

## Use the table

To get information about a postal code, you can get the data using the `Get function`.

    package main

    import (
        "fmt"
        "os
    )

    func main() {
        city, err := PostalCodes.Get(firstPostalCode)
        if err != nil {
            fmt.Printf("An error occurred: %s\n", err)
            os.Exit(1)
        }
        fmt.Printf("The city is called: %s\n", city.City)
    }
