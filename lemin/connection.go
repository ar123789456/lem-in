package lemin

import "strings"

func Connection(list []string) [][]string {
	var splited []string
	var arr [][]string
	for _, v := range list {
		if v[0] == '#' {
			continue
		} else {
			splited = strings.Split(v, "-")
			if len(splited) == 2 {
				arr = append(arr, splited)
			}
		}
	}
	return arr
}
