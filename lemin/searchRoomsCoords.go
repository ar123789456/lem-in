package lemin

import (
	"strconv"
	"strings"
)

func SearchRoomsCoords(list []string) ([][]string, StartEnd, bool) {
	startEnd := StartEnd{}
	var splited []string
	var arr [][]string
	isExist := true
	for i, v := range list {
		if v == "##start" {
			splited = strings.Split(list[i+1], " ")
			if len(splited) == 3 {
				startEnd.Start = splited[0]
			}
		}
		if v == "##end" {
			splited = strings.Split(list[i+1], " ")
			if len(splited) == 3 {
				startEnd.End = splited[0]
			}
		}
		if len(v) == 0 {
			isExist = false
			return arr, startEnd, isExist
		}
		if v[0] == '#' {
			continue
		} else {
			splited = strings.Split(v, " ")
			if len(splited) == 3 {
				arr = append(arr, splited)
			}
			if len(splited) == 1 {
				ant, _ := strconv.Atoi(splited[0])
				if startEnd.Ant == 0 {
					startEnd.Ant = ant
				}
				if startEnd.Ant == 0 || ant < 0 {
					isExist = false
					return arr, startEnd, isExist
				}
			}
		}
	}
	return arr, startEnd, isExist
}
