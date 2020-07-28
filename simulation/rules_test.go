package simulation

import "testing"

func TestNumNeighbors(t *testing.T) {
	val := byte(0b000_01_111)
	if NumNeighbors(val) != 4 {
		t.Fail()
	}
	val = byte(0b111_10_000)
	if NumNeighbors(val) != 4 {
		t.Fail()
	}

	val = byte(0b000_00_000)
	if NumNeighbors(val) != 0 {
		t.Fail()
	}

	val = byte(0b111_11_111)
	if NumNeighbors(val) != 8 {
		t.Fail()
	}

	val = byte(0b0101_0101)
	if NumNeighbors(val) != 4 {
		t.Fail()
	}
}
