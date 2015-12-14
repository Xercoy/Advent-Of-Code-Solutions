package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		presentsList       []Present
		wrappingPaperTotal float64
		ribbonTotal        float64
		verboseFlag        bool
	)

	flag.BoolVar(&verboseFlag, "verbose", false, "print log to stdout.")

	// Read newline separated dimensions, ex. 4x5x11
	dimensionsContent, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err.Error())
	}

	dimensionsList := strings.Split(strings.Trim(string(dimensionsContent),
		"\n"), "\n")

	for _, d := range dimensionsList {
		p := *NewPresent(ParseDimension(d))
		presentsList = append(presentsList, p)

		wrappingPaperTotal += p.SurfaceArea() + p.SmallestSideArea()
		ribbonTotal += p.Volume() + p.SmallestPerimeter()
	}

	fmt.Printf("\n%f feet of wrapping paper is needed.\n", wrappingPaperTotal)

	fmt.Printf("\n%f feet of ribbon is needed.\n", ribbonTotal)

	return
}
