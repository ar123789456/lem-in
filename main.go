package main

import (
	"fmt"
	"lem-in/lemin"
	"os"
)

func main() {
	//Finish Him!!!
	// defer func ()  {

	// }
	args := os.Args[1:]
	lemin.CheckArgs(args)
	file, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	fileInfo := lemin.ReadyFile(file)
	roomsAndCoords, startEnd, isExist := lemin.SearchRoomsCoords(fileInfo)
	lemin.IsExist(isExist, startEnd)
	buildedRooms, startEnd := lemin.BuildRooms(roomsAndCoords, startEnd)
	connect := lemin.Connection(fileInfo)
	buildedRooms = lemin.AddNeighboards(connect, buildedRooms)
	// fmt.Println("Hello")
	pathsList := lemin.OptimalRoads(buildedRooms, startEnd)
	lemin.IsPaths(pathsList)
	reversedPathsList := lemin.ListReverse(pathsList)
	tic := lemin.TicFind(reversedPathsList, startEnd)
	antsStructs := lemin.CreateStructAntsMoving(reversedPathsList, tic, startEnd)
	// fmt.Println(pathsList)
	fmt.Print(string(file))
	lemin.PrintAntsMoving(antsStructs, 0)
	fmt.Println(len(pathsList))
	// time.Sleep(20 * time.Second)
}
