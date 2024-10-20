package conversor

import (
	"strconv"
	"time"
)

type StringConversorInterface interface {
	ToUint64(str string) (uint64, error)
	ToInt(str string) (int, error)
	ToDate(str string) (time.Time, error)
	ToDateTime(str string) (time.Time, error)
	ToDateTimeTZ(str string) (time.Time, error)
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

func (c *stringConversor) ToDate(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}

func (c *stringConversor) ToDateTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05", str)
}

func (c *stringConversor) ToDateTimeTZ(str string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05-07:00", str)
}
