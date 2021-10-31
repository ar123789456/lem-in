package lemin

func AddRoom(maps [][]Honeycomb, graph map[string]Graph) {
	for _, gr := range graph {
		maps[gr.Coor[1]][gr.Coor[0]].Name = gr.RoomName
	}
}
