package simulation

import (
	"testing"
)

func TestDataGrid_NeighborsValue(t *testing.T) {
	grid := make(DataGrid, 3)
	grid[0] = make([]byte, 3)
	grid[1] = make([]byte, 3)
	grid[2] = make([]byte, 3)

	grid[0][0] = ALIVE
	if grid.NeighborsValue(1, 1) != 0b10000000 {
		t.Fail()
	}

	grid[2][2] = ALIVE
	if grid.NeighborsValue(1, 1) != 0b10000001 {
		t.Fail()
	}

	grid[0][2] = ALIVE
	if grid.NeighborsValue(1, 1) != 0b10100001 {
		t.Fail()
	}

	grid[2][0] = ALIVE
	if grid.NeighborsValue(1, 1) != 0b10100101 {
		t.Fail()
	}

	grid[0][1] = ALIVE
	if grid.NeighborsValue(1, 1) != 0b11100101 {
		t.Fail()
	}

	grid[2][1] = ALIVE
	if grid.NeighborsValue(1, 1) != 0b11100111 {
		t.Fail()
	}

	grid[1][0] = ALIVE
	if grid.NeighborsValue(1, 1) != 0b11110111 {
		t.Fail()
	}

	grid[1][2] = ALIVE
	if grid.NeighborsValue(1, 1) != 0b11111111 {
		t.Fail()
	}

}
