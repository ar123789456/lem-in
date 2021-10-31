package lemin

import "fmt"

func CreateStructAntsMoving(reversedPathsList [][]string, tic int, startEnd StartEnd) []*AntsDirection {
	fmt.Println("---------------------")
	ants := startEnd.Ant
	var antsStruct []*AntsDirection
	for i := 0; i < len(reversedPathsList); i++ {
		// fmt.Println("tic", tic, len(reversedPathsList[i]))
		ant := tic - len(reversedPathsList[i])
		ants = ants - ant
		antsDirection := AntsDirection{
			ants:         ant,
			paths:        reversedPathsList[i],
			antsTicQueue: []string{},
		}
		if ants < 0 {
			antsDirection.ants = antsDirection.ants + ants
		}
		antsDirection = antsDirection.EmptyQueue()
		antsStruct = append(antsStruct, &antsDirection)
		fmt.Println(antsDirection.ants, ants, tic, startEnd.Ant)
	}
	return antsStruct
}
