package lemin

func IsContain(g *Honeycomb, listg []*Honeycomb) bool {
	for _, v := range listg {
		if v.Coor == g.Coor {
			return true
		}
	}
	return false
}