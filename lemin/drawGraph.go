package lemin

func DrawGraph(x, y int) [][]Honeycomb {
	var d [][]Honeycomb
	for i := 0; i <= y; i++ {
		var h []Honeycomb
		for j := 0; j <= x; j++ {
			var bee Honeycomb
			bee.Name = " "
			bee.Coor = [2]int{i, j}
			h = append(h, bee)
		}
		d = append(d, h)
	}
	return d
}
