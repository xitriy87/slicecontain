package slicecontain

import (
	"fmt"
)

func Contain(s []string,  x string) bool {
	for _, v := range s {
		if x == v {
			return true
		}
	
	}
	return false
}

func ContainInt(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}


