// Package strconv formats primitive types to strings.
package strconv

// FormatSliceToCSV converts a slice of strings to a Comma Separated Values
// string.
func FormatSliceToCSV(s []string) string {
	var csvString string

	for i, v := range s {
		if i > 0 {
			csvString += ","
		}

		csvString += v
	}

	return csvString
}
