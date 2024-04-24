package strings

// SliceToCSV converts a slice of strings to a CSV string.
func SliceToCSV(s []string) string {
	var csvString string

	for i, v := range s {
		if i > 0 {
			csvString += ","
		}

		csvString += v
	}

	return csvString
}
