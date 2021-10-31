package lemin

func Neighboards(x, y int, maps [][]Honeycomb) [][]Honeycomb {
	for x1 := x - 1; x1 <= x+1; x1++ {
		for y1 := y - 1; y1 <= y+1; y1++ {
			if x1 == x && y1 == y {
				continue
			}
			if x1 < 0 || y1 < 0 || x1 == len(maps) || y1 == len(maps[0]) {
				continue
			}
			maps[x][y].Next = append(maps[x][y].Next, &maps[x1][y1])
		}
	}
	return maps
}
