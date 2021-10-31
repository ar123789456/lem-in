package lemin

func AStar(start, goal Honeycomb) *Honeycomb {
	var closed []*Honeycomb
	var open []*Honeycomb
	start.Begin = 0
	start.End = Orel(&goal, &start)
	open = append(open, &start)
	for len(open) != 0 {
		open = BubleSort(open)
		p := open[0]
		open = open[1:]
		if IsContain(p, closed) {
			continue
		}
		if p.Coor == goal.Coor {
			return p
		}
		closed = append(closed, p)
		for _, v := range p.Next {
			if IsContain(v, closed) {
				continue
			}
			if v.Name != " " && v.Name != goal.Name {
				continue
			}
			v.Begin = Orel(v, p) + p.Begin
			v.End = Orel(&goal, v)
			v.Parent = p
			open = append(open, v)
		}
	}
	return nil
}