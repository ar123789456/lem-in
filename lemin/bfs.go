package lemin

func BFS(startEnd StartEnd) bool {
	var closed map[string]*Rooms
	closed = make(map[string]*Rooms)
	var open map[string]*Rooms
	open = make(map[string]*Rooms)
	open[startEnd.StartGr.name] = startEnd.StartGr

	for len(open) != 0 {
		var now *Rooms
		for k, t := range open {
			now = t
			delete(open, k)
			break
		}
		// if len(startEnd.EndGr.parent) == len(startEnd.EndGr.neighboards) {
		// 	return true
		// }
		if now.name == startEnd.EndGr.name {
			return true
		}
		if Attend(closed, now) {
			continue
		}
		if Attend(open, now) {
			continue
		}
		for _, room := range now.neighboards {
			if Attend(closed, room) {
				continue
			}
			if Attend(open, room) {
				continue
			}
			if room.name == startEnd.EndGr.name {
				room.parent = append(room.parent, now)
				room.singParent = now
				return true
			}
			if room.untouch {
				continue
			}

			room.singParent = now
			open[room.name] = room
		}
		if len(open) == 0 {
			closed[now.name] = now
			for _, room := range now.neighboards {
				if Attend(closed, room) {
					// closed[room.name] = room
					continue
				}
				if room.untouch {
					if room.singParent.name != now.name {
						tem := room.singParent
						// fmt.Println(tem.name, tem.neighboards, closed)
						// now.untouch = true
						if !Attend(closed, tem) {
							room.singParent = now
						}
						open[tem.name] = tem
						break
					}
					// continue
				}
			}
		}
		closed[now.name] = now
	}
	return false
}
