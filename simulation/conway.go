package simulation

//Returns a mapping of the neighbor value to the output
func GenerateConwayNeighborsRules() (alive map[byte]bool, dead map[byte]bool) {
	alive = make(map[byte]bool, 256)
	dead = make(map[byte]bool, 256)

	for neighborhood := byte(0); neighborhood <= byte(254); neighborhood++ {
		alive[neighborhood] = ConwayIsNextStageAlive(true, neighborhood)
		dead[neighborhood] = ConwayIsNextStageAlive(false, neighborhood)
	}
	alive[255] = ConwayIsNextStageAlive(true, 255)
	dead[255] = ConwayIsNextStageAlive(false, 255)

	return alive, dead
}

func ConwayIsNextStageAlive(alive bool, neighborhood byte) bool {
	numNeighbors := NumNeighbors(neighborhood)
	if alive && (numNeighbors == 2 || numNeighbors == 3) {
		return true
	} else if !alive && numNeighbors == 3 {
		return true
	} else {
		return false
	}
}
