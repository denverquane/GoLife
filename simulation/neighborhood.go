package simulation

import (
	"github.com/lucasb-eyer/go-colorful"
	"math/bits"
)

//only call with inner coordinates (don't do perimeters)
func (dg DataGrid) InnerNeighborsValue(y, x uint32) byte {
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

func (dg DataGrid) NeighborsColorAverage(y, x uint32, neighbors byte) uint32 {
	col := colorful.Color{
		R: 0,
		G: 0,
		B: 0,
	}
	num := 0

	if neighbors&NW_MASK > 0 {
		cell := dg[y-1][x-1]
		red, green, blue := rgbOfCell(cell)
		col = colorful.Color{
			R: red,
			G: green,
			B: blue,
		}
		num++
	}
	if neighbors&N_MASK > 0 {
		cell := dg[y-1][x]
		red, green, blue := rgbOfCell(cell)
		if num == 0 {
			col = colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
		} else {
			col2 := colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
			col = col.BlendRgb(col2, 0.5)
		}
		num++
	}
	if neighbors&NE_MASK > 0 {
		cell := dg[y-1][x+1]
		red, green, blue := rgbOfCell(cell)
		if num == 0 {
			col = colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
		} else {
			col2 := colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
			col = col.BlendRgb(col2, 0.5)
		}
		num++
	}
	if neighbors&E_MASK > 0 {
		cell := dg[y][x+1]
		red, green, blue := rgbOfCell(cell)
		if num == 0 {
			col = colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
		} else {
			col2 := colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
			col = col.BlendRgb(col2, 0.5)
		}
		num++
	}
	if neighbors&SE_MASK > 0 {
		cell := dg[y+1][x+1]
		red, green, blue := rgbOfCell(cell)
		if num == 0 {
			col = colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
		} else {
			col2 := colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
			col = col.BlendRgb(col2, 0.5)
		}
		num++
	}
	if neighbors&S_MASK > 0 {
		cell := dg[y+1][x]
		red, green, blue := rgbOfCell(cell)
		if num == 0 {
			col = colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
		} else {
			col2 := colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
			col = col.BlendRgb(col2, 0.5)
		}
		num++
	}
	if neighbors&SW_MASK > 0 {
		cell := dg[y+1][x-1]
		red, green, blue := rgbOfCell(cell)
		if num == 0 {
			col = colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
		} else {
			col2 := colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
			col = col.BlendRgb(col2, 0.5)
		}
		num++
	}
	if neighbors&W_MASK > 0 {
		cell := dg[y][x-1]
		red, green, blue := rgbOfCell(cell)
		if num == 0 {
			col = colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
		} else {
			col2 := colorful.Color{
				R: red,
				G: green,
				B: blue,
			}
			col = col.BlendRgb(col2, 0.5)
		}
		num++
	}
	newRed := col.R * 255.0
	newGreen := col.G * 255.0
	newBlue := col.B * 255.0
	return uint32(newRed)<<24 + uint32(newGreen)<<16 + uint32(newBlue)<<8 + ALIVE
}

func rgbOfCell(cell uint32) (r, g, b float64) {
	return float64((cell>>24)&ALIVE) / 255.0, float64((cell>>16)&ALIVE) / 255.0, float64((cell>>8)&ALIVE) / 255.0
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

const (
	NW_MASK byte = 0b0000_0001
	N_MASK  byte = 0b0000_0010
	NE_MASK byte = 0b0000_0100
	E_MASK  byte = 0b0000_1000
	SE_MASK byte = 0b0001_0000
	S_MASK  byte = 0b0010_0000
	SW_MASK byte = 0b0100_0000
	W_MASK  byte = 0b1000_0000
)

//Northwest is the right-most bit of the byte, rotating around clockwise until the left side of the byte
//So the neighborhood can be interpreted as so:
//0b_W_SW_S_SE_E_NE_N_NW

func (dg DataGrid) PerimeterNeighborsValue(dir Direction, y, x uint32) byte {
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
