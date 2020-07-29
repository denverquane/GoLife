package simulation

import (
	"math/bits"
)

//only call with inner coordinates (don't do perimeters)
func (dg DataGrid) InnerNeighborsValue(y, x int64) byte {
	neighborState := byte(0)

	neighborState |= isAlive(dg[y-1][x-1])
	neighborState |= isAlive(dg[y-1][x]) >> 1
	neighborState |= isAlive(dg[y-1][x+1]) >> 2

	neighborState |= isAlive(dg[y][x-1]) >> 3
	neighborState |= isAlive(dg[y][x+1]) >> 4

	neighborState |= isAlive(dg[y+1][x-1]) >> 5
	neighborState |= isAlive(dg[y+1][x]) >> 6
	neighborState |= isAlive(dg[y+1][x+1]) >> 7

	return neighborState
}

type Direction byte

const (
	N  Direction = 1
	NE Direction = 2
	E  Direction = 3
	SE Direction = 4
	S  Direction = 5
	SW Direction = 6
	W  Direction = 7
	NW Direction = 8
)

func (dg DataGrid) PerimeterNeighborsValue(dir Direction, y, x int64) byte {
	neighborState := byte(0)
	switch dir {
	case N:
		neighborState |= isAlive(dg[y][x-1]) >> 3
		neighborState |= isAlive(dg[y][x+1]) >> 4
		neighborState |= isAlive(dg[y+1][x-1]) >> 5
		neighborState |= isAlive(dg[y+1][x]) >> 6
		neighborState |= isAlive(dg[y+1][x+1]) >> 7
	case NE:
		neighborState |= isAlive(dg[y][x-1]) >> 3
		neighborState |= isAlive(dg[y+1][x-1]) >> 5
		neighborState |= isAlive(dg[y+1][x]) >> 6
	case E:
		neighborState |= isAlive(dg[y-1][x-1])
		neighborState |= isAlive(dg[y-1][x]) >> 1
		neighborState |= isAlive(dg[y][x-1]) >> 3
		neighborState |= isAlive(dg[y+1][x-1]) >> 5
		neighborState |= isAlive(dg[y+1][x]) >> 6
	case SE:
		neighborState |= isAlive(dg[y-1][x-1])
		neighborState |= isAlive(dg[y-1][x]) >> 1
		neighborState |= isAlive(dg[y][x-1]) >> 3
	case S:
		neighborState |= isAlive(dg[y-1][x-1])
		neighborState |= isAlive(dg[y-1][x]) >> 1
		neighborState |= isAlive(dg[y-1][x+1]) >> 2
		neighborState |= isAlive(dg[y][x-1]) >> 3
		neighborState |= isAlive(dg[y][x+1]) >> 4
	case SW:
		neighborState |= isAlive(dg[y-1][x]) >> 1
		neighborState |= isAlive(dg[y-1][x+1]) >> 2
		neighborState |= isAlive(dg[y][x+1]) >> 4
	case W:
		neighborState |= isAlive(dg[y-1][x]) >> 1
		neighborState |= isAlive(dg[y-1][x+1]) >> 2
		neighborState |= isAlive(dg[y][x+1]) >> 4
		neighborState |= isAlive(dg[y+1][x]) >> 6
		neighborState |= isAlive(dg[y+1][x+1]) >> 7
	case NW:
		neighborState |= isAlive(dg[y][x+1]) >> 4
		neighborState |= isAlive(dg[y+1][x]) >> 6
		neighborState |= isAlive(dg[y+1][x+1]) >> 7
	}
	return neighborState
}

func isAlive(cell byte) byte {
	return cell & ALIVE
}

func NumNeighbors(neighborhood byte) int {
	return bits.OnesCount(uint(neighborhood))
}
