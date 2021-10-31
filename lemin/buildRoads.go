package lemin

func BuildRoads(maps [][]Honeycomb, connect [][]string, graph map[string]Graph) [][]Honeycomb {
	// fmt.Println(connect)
	for _, v := range connect[:5] /*connect[:len(connect)]*/ {
		astarResult := AStar(maps[graph[v[0]].Coor[1]][graph[v[0]].Coor[0]], maps[graph[v[1]].Coor[1]][graph[v[1]].Coor[0]])
		Draw(astarResult)
		// GraphClear(&maps[0][0])
	}
	return maps
}