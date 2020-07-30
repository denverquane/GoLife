package simulation

import (
	"log"
	"testing"
)

func TestDataGrid_InnerNeighborsValue(t *testing.T) {
	dg := make(DataGrid, 10)
	for y := 0; y < 10; y++ {
		dg[y] = make([]uint32, 10)
	}

	dg[3][4] = ALIVE_FULL
	dg[4][4] = ALIVE_FULL
	dg[4][3] = ALIVE_FULL
	dg[4][2] = ALIVE_FULL

	val := dg.InnerNeighborsValue(3, 3)
	log.Print(ConwayIsNextStageAlive(false, val))
	log.Printf("%08b\n", val)
	if NumNeighbors(val) != 4 {
		t.Fail()
	}
}
