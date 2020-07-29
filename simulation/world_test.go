package simulation

import (
	"log"
	"testing"
	"time"
)

func TestDataGrid_NeighborsValue(t *testing.T) {
	grid := make(DataGrid, 3)
	grid[0] = make([]byte, 3)
	grid[1] = make([]byte, 3)
	grid[2] = make([]byte, 3)

	grid[0][0] = ALIVE
	if grid.InnerNeighborsValue(1, 1) != 0b100_00_000 {
		t.Fail()
	}

	grid[2][2] = ALIVE
	if grid.InnerNeighborsValue(1, 1) != 0b100_00_001 {
		t.Fail()
	}

	grid[0][2] = ALIVE
	if grid.InnerNeighborsValue(1, 1) != 0b101_00_001 {
		t.Fail()
	}

	grid[2][0] = ALIVE
	if grid.InnerNeighborsValue(1, 1) != 0b101_00_101 {
		t.Fail()
	}

	grid[0][1] = ALIVE
	if grid.InnerNeighborsValue(1, 1) != 0b111_00_101 {
		t.Fail()
	}

	grid[2][1] = ALIVE
	if grid.InnerNeighborsValue(1, 1) != 0b111_00_111 {
		t.Fail()
	}

	grid[1][0] = ALIVE
	if grid.InnerNeighborsValue(1, 1) != 0b111_10_111 {
		t.Fail()
	}

	grid[1][2] = ALIVE
	if grid.InnerNeighborsValue(1, 1) != 0b111_11_111 {
		t.Fail()
	}
}

func TestWorld_Tick(t *testing.T) {
	const Iterations = 100
	world := NewConwayWorld(10000, 10000)
	now := time.Now().UnixNano()
	for i := 0; i < Iterations; i++ {
		world.Tick()
	}
	end := time.Now().UnixNano() - now
	log.Printf("Took %fms to complete %d iterations on a 10k^2 grid", float64(end)/1_000_000.0, Iterations)
}

func TestWorld_Tick2(t *testing.T) {
	world := NewConwayWorld(10, 10)
	world.MarkAlive(0, 0)
	world.MarkAlive(9, 9)
	world.MarkAlive(0, 9)
	world.MarkAlive(9, 0)

	world.Tick()

	if (*world.data)[0][0]&ALIVE == ALIVE || (*world.data)[9][9]&ALIVE == ALIVE || (*world.data)[0][9]&ALIVE == ALIVE || (*world.data)[9][0]&ALIVE == ALIVE {
		t.Fail()
	}

	world.MarkAlive(0, 1)
	world.MarkAlive(0, 2)
	world.MarkAlive(0, 3)
	world.Tick()
	if (*world.data)[0][2]&ALIVE != ALIVE || (*world.data)[1][2]&ALIVE != ALIVE {
		t.Fail()
	}
}
