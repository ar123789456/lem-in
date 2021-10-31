package lemin

func Checker(str []string) bool {
	for _, i := range str {
		if i != "" {
			return true
		}
	}
	return false
}
