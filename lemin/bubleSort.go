package lemin

func BubleSort(open []*Honeycomb) []*Honeycomb {
	for i := 0; i < len(open); i++ {
		for j := 0; j < len(open)-1-i; j++ {
			if open[j].Begin+open[j].End > open[j+1].Begin+open[j+1].End {
				open[j], open[j+1] = open[j+1], open[j]
			}
		}
	}
	return open
}