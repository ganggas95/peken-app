package helper

import "strconv"

// ConvertStringToFloat64 convert string to float64
func ConvertStringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return f
}

// ConvertStringToInt convert string to int
func ConvertStringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// ConvertStringToInt64 convert string to int64
func ConvertStringToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// ConvertStringToBool convert string to bool
func ConvertStringToBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return b
}
