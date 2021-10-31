package lemin

func OptimalRoads(buildedRooms map[string]*Rooms, startEnd StartEnd) [][]string {
	roads := 0
	long := 0
	oldfinish := 0
	var ret [][]string
	for len(startEnd.EndGr.neighboards) != roads {
		tic := BFS(startEnd)
		// fmt.Println(ret)
		Indexer(startEnd)
		if tic {
			long = lenList(startEnd.EndGr)
			roads++
		}
		finish := (long + int(startEnd.Ant)) / roads
		if (long+int(startEnd.Ant))%roads != 0 {
			finish++
		}
		// fmt.Println(finish, long)
		if !tic {
			// return
			return ret
		}
		if finish > oldfinish && oldfinish != 0 {
			// return
			return ret
		}
		if int(startEnd.Ant) < roads {
			return ret
		}
		ret = CreatePathsList(startEnd.EndGr)
		if len(ret[0]) == 1 {
			val := ret[0]
			for i := len(ret); i < startEnd.Ant; i++ {
				ret = append(ret, val)
			}
			return ret
		}
		oldfinish = finish
	}
	return ret
}

func lenList(end *Rooms) int {
	a := 0
	for _, v := range end.parent {
		a++
		for v.singParent != nil {
			a++
			v = v.singParent
		}
	}
	return a
}
