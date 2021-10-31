package lemin

func CreatePathsList(end *Rooms) [][]string {
	var names [][]string
	for _, v := range end.parent {
		var arr []string
		arr = append(arr, end.name)
		for v.singParent != nil {
			arr = append(arr, v.name)
			v = v.singParent
		}
		names = append(names, arr)
	}
	return names
}
