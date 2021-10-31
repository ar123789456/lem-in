package lemin

func GraphClear(gr *Honeycomb) {
	var closed []*Honeycomb
	var open []*Honeycomb
	open = append(open, gr)
	for len(open) != 0 {
		p := open[0]
		open = open[1:]
		p.Begin = 0
		p.End = 0
		p.Parent = nil
		if IsContain(p, closed) {
			continue
		}
		closed = append(closed, p)
		for _, v := range p.Next {
			if IsContain(v, closed) {
				continue
			}
			open = append(open, v)
		}
	}
}