package lemin

func TicFind(paths [][]string, startEnd StartEnd) int {
	antscount := startEnd.Ant
	sumOfLenght := 0
	countOfPaths := 0
	tic := 0
	tmp := 0
	for _, v := range paths {
		sumOfLenght += len(v)
		countOfPaths++
	}
	tmp = sumOfLenght + antscount
	tic = tmp / countOfPaths
	if tmp%countOfPaths != 0 {
		tic++
	}
	return tic
}
