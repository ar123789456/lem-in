package lemin

type StartEnd struct {
	Ant     int
	Start   string
	StartGr *Rooms
	End     string
	EndGr   *Rooms
}

type Rooms struct {
	name        string
	status      string
	untouch     bool
	parent      []*Rooms
	singParent  *Rooms
	neighboards []*Rooms
}

type Honeycomb struct {
	Name   string
	Next   []*Honeycomb
	Parent *Honeycomb
	Coor   [2]int
	Begin  float64
	End    float64
}

type Graph struct {
	Coor     [2]int
	RoomName string
}

type AntsDirection struct {
	ants         int
	paths        []string
	antsTicQueue []string
}
