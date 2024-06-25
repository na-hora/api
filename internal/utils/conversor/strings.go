package conversor

import (
	"strconv"
)

type StringConversorInterface interface {
	ToUint64(str string) (uint64, error)
}

type stringConversor struct{}

func GetStringConversor() StringConversorInterface {
	return &stringConversor{}
}

func (c *stringConversor) ToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
