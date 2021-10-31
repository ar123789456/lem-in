package lemin

func ListReverse(p [][]string) [][]string {
	for i, v := range p {
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
		}
		p[i] = v
	}
	return p
}
