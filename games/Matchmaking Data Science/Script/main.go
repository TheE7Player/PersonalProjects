package main

import (
	f "CSGO_STATISTICS/filter"
	d "CSGO_STATISTICS/setup"
	"fmt"
)

// Global variables

// Will hold the table structure after reading in the csv file
var table map[string][]string

// Entry-point of the script
func main() {
	// Load the file to read
	data := d.Setup("./csgo_data.csv")

	fmt.Println("Data has been assigned, transforming...")

	// Load the file into a map structure for fast validation
	table = d.Setup_table(data)

	// Now we do the science (statistics finding)
	f.Run(table)
}
