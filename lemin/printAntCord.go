package lemin

import "fmt"

func PrintAntCord(antsDirection []*AntsDirection) {
	for _, v := range antsDirection {
		for j := len(v.antsTicQueue) - 1; j >= 0; j-- {
			if v.antsTicQueue[j] != "" {
				fmt.Print(v.antsTicQueue[j], "-", v.paths[j], " ")
			}
		}
	}

}
