package lib

import (
	"strconv"
	"strings"
)

func UIntPlease(s string) uint64 {
	ret, _ := strconv.ParseUint(s, 10, 0)
	return ret
}

func UIntsPlease(s, sep string) []uint64 {
	var ret []uint64
	for _, el := range strings.Split(s, sep) {
		trimmed := strings.TrimSpace(el)
		if trimmed != "" {
			ret = append(ret, UIntPlease(trimmed))
		}
	}
	return ret
}

func IntPlease(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func IntsPlease(s, sep string) []int {
	var ret []int
	for _, el := range strings.Split(s, sep) {
		trimmed := strings.TrimSpace(el)
		if trimmed != "" {
			ret = append(ret, IntPlease(trimmed))
		}
	}
	return ret
}
