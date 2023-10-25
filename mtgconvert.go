// +build js,wasm

package main

import (
	"encoding/csv"
	"strings"
	"syscall/js"
	"time"
)

func convertCSV(this js.Value, p []js.Value) interface{} {
	inputCSV := p[0].String()
	isFirstType := strings.HasPrefix(inputCSV, "sep=,")
	reader := csv.NewReader(strings.NewReader(inputCSV))

	// Remove "sep=," line if it exists
	if isFirstType {
		reader.Read()
	}

	// Read the header line
	reader.Read()

	var outputCSV string
	outputCSV += "\"Count\",\"Tradelist Count\",\"Name\",\"Edition\",\"Condition\",\"Language\",\"Foil\",\"Tags\",\"Last Modified\",\"Collector Number\",\"Alter\",\"Proxy\",\"Purchase Price\"\n"

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		// Assume processFirstType for simplicity
		newRow := []string{}
		// Simplified logic for the example
		newRow = append(newRow, row[0], row[1], row[2], strings.ToLower(row[3]), "Near Mint", row[8], "", "", time.Now().Format("2006-01-02 15:04:05.000000"), row[6], "False", "False", row[10])
		outputCSV += "\"" + strings.Join(newRow, "\",\"") + "\"\n"
	}

	return outputCSV
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("convertCSV", js.FuncOf(convertCSV))
	<-c
}
