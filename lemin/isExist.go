package lemin

func IsExist(isExist bool, startEnd StartEnd) {
	if !isExist {
		PrintError()
	}
	if len(startEnd.Start) == 0 || len(startEnd.End) == 0 {
		PrintError()
	}
}
