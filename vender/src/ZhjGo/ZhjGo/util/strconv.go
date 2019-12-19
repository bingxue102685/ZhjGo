package util

import (
	"ZhjGo/ZhjGo/log"
	"strconv"
)

func Str2int_ZhjGo(data string) int {
	v, err := strconv.Atoi(data)
	if err != nil {
		log.Debug(err.Error())
		return -1
	}
	return v
}

func Int2str_ZhjGo(data int) string {
	return strconv.Itoa(data)
}
