package utils

import (
	"strconv"
)

// String2Int Convert a string to an int
func String2Int(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}

// Int2String Convert an int to a string
func Int2String(i int) string {
	return strconv.Itoa(i)
}
