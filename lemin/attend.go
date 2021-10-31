package lemin

func AttendOpen(closed []*Rooms, now *Rooms) bool {
	for _, cl := range closed {
		if cl.name == now.name {
			return true
		}
	}
	return false
}

func Attend(closed map[string]*Rooms, now *Rooms) bool {
	if closed[now.name] != nil {
		return true
	}
	return false
}
