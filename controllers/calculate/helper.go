package calculate

import "strconv"

func formatFloat(value float64, precision int) string {
	return strconv.FormatFloat(value, 'g', precision, 64)
}
