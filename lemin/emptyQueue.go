package lemin

func (antsDirection AntsDirection) EmptyQueue() AntsDirection {
	for i := 0; i < len(antsDirection.paths); i++ {
		antsDirection.antsTicQueue = append(antsDirection.antsTicQueue, "")
	}
	return antsDirection
}
