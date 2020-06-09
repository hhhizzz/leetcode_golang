package _165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < len(v1) && i < len(v2); i++ {
		a, _ := strconv.Atoi(v1[i])
		b, _ := strconv.Atoi(v2[i])
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
	}
	if len(v1) == len(v2) {
		return 0
	} else if len(v1) > len(v2) {
		sum := 0
		for i := len(v2); i < len(v1); i++ {
			s, _ := strconv.Atoi(v1[i])
			sum += s
		}
		if sum == 0 {
			return 0
		} else {
			return 1
		}
	} else {
		sum := 0
		for i := len(v1); i < len(v2); i++ {
			s, _ := strconv.Atoi(v2[i])
			sum += s
		}
		if sum == 0 {
			return 0
		} else {
			return -1
		}
	}
}
