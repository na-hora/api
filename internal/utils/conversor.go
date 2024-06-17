package utils

import (
	"strconv"
)

func StringToUint(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
