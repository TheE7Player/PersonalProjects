package setup

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func Setup(file string) [][]string {
	// Open File
	csvFile, err := os.Open(file)

	// Show an error to console if an error has occured (err is nil, if not)
	if err != nil {
		fmt.Println("Error happen will reading csv: ", err)
	}

	// Create the reader object using the file as IO stream
	reader := csv.NewReader(csvFile)

	// Get the data into a variable called 'data', ignoring the error variable using underscore '_'
	data, _ := reader.ReadAll()

	// Finally, we return back the data to the caller
	return data
}

func Setup_table(file [][]string) map[string][]string {

	// Assign the column length
	columnCount := len(file[0])
	rowSize := len(file)

	fmt.Println("Attempting to read rows of size:", (rowSize - 1))
	// Create the shorthand variable of a map of string arrays ( with fixed length -> columnCount )
	returnMap := make(map[string][]string, columnCount)

	// Create the keys and set the array length with constant
	// First row is the headers, then we allocate new string arrays of fixed size ( rowSize )

	// Create a map with a fixed size which holds the column name based on passed in index
	colNum := make(map[int]string, columnCount)

	for i := 0; i < columnCount; i++ {
		// Create a column name with a fixed string array of row length
		returnMap[file[0][i]] = make([]string, rowSize)

		// Then create a column name index location, to allow the program to insert the right information into the right key.
		colNum[i] = file[0][i]
	}

	// Now we fill in each row (Starting at 1, skipping headers! )

	rSize := rowSize - 1 // Cached calculation result of the amount of iterations required (faster)
	colName := ""
	value := ""
	for i := 1; i < rSize; i++ {
		for k := 0; k < columnCount; k++ {

			// Get the corrisponding key name based on the index we're using
			colName = colNum[k]

			// Get the value of the index of the 2D we're accessing
			value = file[i][k]

			// Validate if it contains any gap in the text (index > -1, if so)
			if strings.Index(value, " ") > -1 {
				value = strings.TrimSpace(value)
			}

			// Set the value based on a key, to the corrisponding index of the key value
			returnMap[colName][i-1] = value
		}
	}

	// Finally return back the structure back to the user
	return returnMap
}
