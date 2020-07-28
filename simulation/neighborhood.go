package simulation

import "math/bits"

func (dg DataGrid) NeighborsValue(y, x int64) byte {
	neighborState := byte(0)
	yMax := int64(len(dg) - 1)
	xMax := int64(len(dg[0]) - 1)
	//cell := dg[y][x]

	if y > 0 {
		if x > 0 {
			neighborState |= isAlive(dg[y-1][x-1])
		}
		neighborState |= isAlive(dg[y-1][x]) >> 1
		if x < xMax {
			neighborState |= isAlive(dg[y-1][x+1]) >> 2
		}
	}

	if x > 0 {
		neighborState |= isAlive(dg[y][x-1]) >> 3
	}

	if x < xMax {
		neighborState |= isAlive(dg[y][x+1]) >> 4
	}

	if y < yMax {
		if x > 0 {
			neighborState |= isAlive(dg[y+1][x-1]) >> 5
		}
		neighborState |= isAlive(dg[y+1][x]) >> 6
		if x < xMax {
			neighborState |= isAlive(dg[y+1][x+1]) >> 7
		}
	}

	return neighborState
}

func isAlive(cell byte) byte {
	return cell & ALIVE
}

func NumNeighbors(neighborhood byte) int {
	return bits.OnesCount(uint(neighborhood))
}
