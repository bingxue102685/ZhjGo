package Mysyl_DB

import (
	"ZhjGo/ZhjGo/util"
)

/*
   数据库的通用函数
*/

func CheckErr(err error) {
	util.CheckErr_ZhjGo(err)
}

func GetTime() string {
	return util.GetTime_ZhjGo()
}

func GetNowtimeMD5() string {
	return util.GetNowtimeMD5_ZhjGo()
}

func GetMD5Hash(text string) string {
	return util.MD5_ZhjGo(text)
}
