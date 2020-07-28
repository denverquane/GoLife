package simulation

//Returns a mapping of the neighbor value to the output
func GenerateConwayNeighborsRules() (alive map[byte]byte, dead map[byte]byte) {
	alive = make(map[byte]byte, 256)
	dead = make(map[byte]byte, 256)

	for neighborhood := byte(0); neighborhood <= byte(254); neighborhood++ {
		alive[neighborhood] = ConwayNewState(ALIVE, neighborhood)
		dead[neighborhood] = ConwayNewState(DEAD, neighborhood)
	}
	alive[255] = ConwayNewState(ALIVE, 255)
	dead[255] = ConwayNewState(DEAD, 255)

	return alive, dead
}

func ConwayNewState(alive byte, neighborhood byte) byte {
	aliveBool := alive == ALIVE
	numNeighbors := NumNeighbors(neighborhood)
	if aliveBool && (numNeighbors == 2 || numNeighbors == 3) {
		return ALIVE
	} else if !aliveBool && numNeighbors == 3 {
		return ALIVE
	} else {
		return DEAD
	}
}
