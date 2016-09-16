/*
I created this file, because I couldn't find an example where an actual
CSV file was being opened and read.  All the examples I found created
the contents of a CSV file programmatically into a variable, which is
not realistic.
*/
package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Constants
const csvFile string = "data.csv" // assume header row is longest row

func main() {
	body, err := ioutil.ReadFile(csvFile)
	if err != nil {
		log.Fatal("ioutil.ReadFile():", err)
	}
	r := csv.NewReader(strings.NewReader(string(body)))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("ReadAll():", err)
	}
	numRows := len(records)
	numCols := len(records[0]) // assumes 1st row is header row
	fmt.Println("# of rows =", numRows)
	fmt.Println("# of columns =", numCols)
	for y, rv := range records {
		// fmt.Println(rv)  // prints out the row
		for x, cv := range rv {
			fmt.Printf("(%d,%d) %s\n", x, y, cv) // prints out each cell w/ x,y
		}
	}
}
