package lemin

func ReverseString(incoming []string) []string {
	// fmt.Println(incoming)
	var arr []string
	for i := len(incoming) - 1; i >= 0; i-- {
		arr = append(arr, incoming[i])
	}
	return arr
}
