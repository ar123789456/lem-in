package lemin

func BuildRooms(roomsAndCoords [][]string, startEnd StartEnd) (map[string]*Rooms, StartEnd) {
	var rooms map[string]*Rooms
	rooms = make(map[string]*Rooms)
	for _, v := range roomsAndCoords {
		var room Rooms
		room.name = v[0]
		if room.name == startEnd.Start {
			room.status = "start"
			startEnd.StartGr = &room
		}
		if room.name == startEnd.End {
			room.status = "end"
			startEnd.EndGr = &room
		}
		rooms[room.name] = &room
		// fmt.Println(rooms[room.name], " rooms")
	}
	return rooms, startEnd
}
