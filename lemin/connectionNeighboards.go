package lemin

func ConnectionNeighboards(maps [][]Honeycomb) [][]Honeycomb {
	for v1 := 0; v1 < len(maps); v1++ {
		for v2 := 0; v2 < len(maps[v1]); v2++ {
			maps = Neighboards(v1, v2, maps)
		}
	}
	return maps
}