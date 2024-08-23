package conversor

import (
	"strconv"
)

type StringConversorInterface interface {
	ToUint64(str string) (uint64, error)
	ToInt(str string) (int, error)
}

type stringConversor struct{}

func GetStringConversor() StringConversorInterface {
	return &stringConversor{}
}

func (c *stringConversor) ToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func (c *stringConversor) ToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
