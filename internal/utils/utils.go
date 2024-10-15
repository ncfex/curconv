package utils

import (
	"fmt"
	"strconv"
)

func FormatFloat(f float64, decimals int) string {
	return fmt.Sprintf("%."+strconv.Itoa(decimals)+"f", f)
}

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
