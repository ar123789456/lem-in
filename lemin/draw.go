package lemin

func Draw(astarResult *Honeycomb) {
	for astarResult != nil {
		//fmt.Println(astarResult.Coor, astarResult.Begin+astarResult.End)
		// if astarResult.Begin
		if astarResult.Name == " " {
			astarResult.Name = "."
		}
		astarResult = astarResult.Parent
	}
}