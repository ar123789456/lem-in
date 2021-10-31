package lemin

func AddNeighboards(connect [][]string, buildedRooms map[string]*Rooms) map[string]*Rooms {
	for _, v := range connect {
		tmp0 := buildedRooms[v[0]]
		tmp1 := buildedRooms[v[1]]
		tmp0.neighboards = append(tmp0.neighboards, tmp1)
		tmp1.neighboards = append(tmp1.neighboards, tmp0)
		buildedRooms[v[0]] = tmp0
		buildedRooms[v[1]] = tmp1
	}
	return buildedRooms
}