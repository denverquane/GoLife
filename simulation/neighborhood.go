package simulation

import (
	"math/bits"
)

//only call with inner coordinates (don't do perimeters)
func (dg DataGrid) InnerNeighborsValue(y, x int64) byte {
	neighborState := byte(0)

	neighborState |= isAlive(dg[y-1][x-1])      //NW
	neighborState |= isAlive(dg[y-1][x]) << 1   //N
	neighborState |= isAlive(dg[y-1][x+1]) << 2 //NE
	neighborState |= isAlive(dg[y][x+1]) << 3   //E
	neighborState |= isAlive(dg[y+1][x+1]) << 4 //SE
	neighborState |= isAlive(dg[y+1][x]) << 5   //S
	neighborState |= isAlive(dg[y+1][x-1]) << 6 //SW
	neighborState |= isAlive(dg[y][x-1]) << 7   //W

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

//Northwest is the right-most bit of the byte, rotating around clockwise until the left side of the byte
//So the neighborhood can be interpreted as so:
//0b_W_SW_S_SE_E_NE_N_NW

func (dg DataGrid) PerimeterNeighborsValue(dir Direction, y, x int64) byte {
	neighborState := byte(0)
	switch dir {
	case N:
		neighborState |= isAlive(dg[y][x+1]) << 3   //E
		neighborState |= isAlive(dg[y+1][x+1]) << 4 //SE
		neighborState |= isAlive(dg[y+1][x]) << 5   //S
		neighborState |= isAlive(dg[y+1][x-1]) << 6 //SW
		neighborState |= isAlive(dg[y][x-1]) << 7   //W
	case NE:
		neighborState |= isAlive(dg[y+1][x]) << 5   //S
		neighborState |= isAlive(dg[y+1][x-1]) << 6 //SW
		neighborState |= isAlive(dg[y][x-1]) << 7   //W
	case E:
		neighborState |= isAlive(dg[y-1][x-1])      //NW
		neighborState |= isAlive(dg[y-1][x]) << 1   //N
		neighborState |= isAlive(dg[y+1][x]) << 5   //S
		neighborState |= isAlive(dg[y+1][x-1]) << 6 //SW
		neighborState |= isAlive(dg[y][x-1]) << 7   //W
	case SE:
		neighborState |= isAlive(dg[y-1][x-1])    //NW
		neighborState |= isAlive(dg[y-1][x]) << 1 //N
		neighborState |= isAlive(dg[y][x-1]) << 7 //W
	case S:
		neighborState |= isAlive(dg[y-1][x-1])      //NW
		neighborState |= isAlive(dg[y-1][x]) << 1   //N
		neighborState |= isAlive(dg[y-1][x+1]) << 2 //NE
		neighborState |= isAlive(dg[y][x+1]) << 3   //E
		neighborState |= isAlive(dg[y][x-1]) << 7   //W
	case SW:
		neighborState |= isAlive(dg[y-1][x]) << 1   //N
		neighborState |= isAlive(dg[y-1][x+1]) << 2 //NE
		neighborState |= isAlive(dg[y][x+1]) << 3   //E
	case W:
		neighborState |= isAlive(dg[y-1][x]) << 1   //N
		neighborState |= isAlive(dg[y-1][x+1]) << 2 //NE
		neighborState |= isAlive(dg[y][x+1]) << 3   //E
		neighborState |= isAlive(dg[y+1][x+1]) << 4 //SE
		neighborState |= isAlive(dg[y+1][x]) << 5   //S
	case NW:
		neighborState |= isAlive(dg[y][x+1]) << 3   //E
		neighborState |= isAlive(dg[y+1][x+1]) << 4 //SE
		neighborState |= isAlive(dg[y+1][x]) << 5   //S
	}
	return neighborState
}

func isAlive(cell uint32) byte {
	if cell&ALIVE > 0 {
		return byte(1)
	} else {
		return byte(0)
	}
}

func isAliveBool(cell uint32) bool {
	return cell&ALIVE > 0
}

func NumNeighbors(neighborhood byte) int {
	return bits.OnesCount(uint(neighborhood))
}
