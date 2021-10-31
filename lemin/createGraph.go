package lemin

import "strconv"

func CreateGraph(roomsAndCoords [][]string) (map[string]Graph, int, int) {
	var arrGraphs map[string]Graph
	arrGraphs = make(map[string]Graph)
	var graph Graph
	x := 0
	y := 0
	for _, v := range roomsAndCoords {
		graph.RoomName = v[0]
		a, _ := strconv.Atoi(v[1])
		b, _ := strconv.Atoi(v[2])
		graph.Coor[0] = a
		graph.Coor[1] = b
		arrGraphs[graph.RoomName] = graph
		if x < a {
			x = a
		}
		if y < b {
			y = b
		}
	}
	return arrGraphs, (x * 3) / 2, (y * 3) / 2
}
