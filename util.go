package main

import (
	"strings"
)

func trimSpace(ss []string) []string {
	ret := make([]string, 0, len(ss))
	for _, s := range ss {
		s = strings.TrimSpace(s)
		if len(s) > 0 {
			ret = append(ret, s)
		}
	}
	return ret
}
