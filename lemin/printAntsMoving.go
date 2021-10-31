package lemin

import (
	"fmt"
	"strconv"
)

func PrintAntsMoving(antsDirection []*AntsDirection, n int) {
	check := 0
	for _, v := range antsDirection {
		if v.ants == 0 && !Checker(v.antsTicQueue) {
			check++
		}
		if check == len(antsDirection) {
			// fmt.Println("\t")
			return
		}
		for j := len(v.antsTicQueue) - 1; j >= 0; j-- {
			if j == 0 {
				if v.ants != 0 {
					v.ants--
					n++
					v.antsTicQueue[j] = "L" + strconv.Itoa(n)
				} else {
					v.antsTicQueue[j] = ""
				}
				continue
			}
			v.antsTicQueue[j] = v.antsTicQueue[j-1]
		}
	}
	fmt.Print("\n")
	PrintAntCord(antsDirection)
	PrintAntsMoving(antsDirection, n)
}
