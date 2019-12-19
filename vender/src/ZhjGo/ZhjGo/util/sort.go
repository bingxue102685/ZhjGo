package util

import (
	"ZhjGo/ZhjGo/conf"
)

func Sort_ZhjGo(data map[string]*conf.DSQ_Exp, iExp int) int {

	if iExp == 0 {
		return 0
	}
	var length = len(data)
	var ssort []int

	for _, v := range data {
		ssort = append(ssort, Str2int_ZhjGo(v.Exp))
	}

	for i := 1; i < length; i++ {
		for j := i; j > 0 && ssort[j] < ssort[j-1]; j-- {
			ssort[j], ssort[j-1] = ssort[j-1], ssort[j]
		}
	}
	for index, val := range ssort {
		if iExp == val {
			return index
		}
	}
	return 0
}
