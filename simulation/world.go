package simulation

type DataGrid [][]byte

const ALIVE byte = 0b10000000

type Grid struct {
	width  int64
	height int64
	//indexed [y][x]!
	data         DataGrid
	aliveMapping map[byte]byte
	deadMapping  map[byte]byte
}

func NewConwayWorld(width, height int64) Grid {
	data := make(DataGrid, height)
	for i, _ := range data {
		data[i] = make([]byte, width)
	}
	return Grid{
		width:  width,
		height: height,
		data:   data,
	}
}

func (dg DataGrid) NeighborsValue(y, x int64) byte {
	neighborState := byte(0)
	//cell := dg[y][x]

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

func isAlive(cell byte) byte {
	return cell & ALIVE
}
