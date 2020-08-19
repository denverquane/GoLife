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

//only gets called with 3 neighbors
func (dg *DataGrid) NewCellNeighborsColorBlend(y, x uint32, neighbors byte) uint32 {
	retColor := colorful.Color{}
	colorsBlended := 0
	for i := N; i < 9; i++ {
		if neighbors&DirectionMasks[i] > 0 {
			xOff := XOffsets[i]
			yOff := YOffsets[i]
			cell := (*dg)[int(y)+yOff][int(x)+xOff]
			col := colorOfCell(cell)
			if colorsBlended == 0 {
				retColor = col
			} else if colorsBlended == 1 {
				retColor = retColor.BlendRgb(col, 0.5)
			} else if colorsBlended == 2 {
				retColor = retColor.BlendRgb(col, 0.333)
			}
			colorsBlended++
		}
	}
	newRed := retColor.R * 255.0
	newGreen := retColor.G * 255.0
	newBlue := retColor.B * 255.0
	return uint32(newRed)<<24 + uint32(newGreen)<<16 + uint32(newBlue)<<8 + ALIVE
}

func (dg *DataGrid) ExistingCellNeighborsColorBlend(oldCell uint32, y, x uint32, neighbors byte) uint32 {
	retColor := colorOfCell(oldCell)
	colorsBlended := 0
	for i := N; i < 9; i++ {
		if neighbors&DirectionMasks[i] > 0 {
			xOff := XOffsets[i]
			yOff := YOffsets[i]
			cell := (*dg)[int(y)+yOff][int(x)+xOff]
			col := colorOfCell(cell)
			if colorsBlended == 0 {
				retColor = retColor.BlendRgb(col, 0.5)
			} else if colorsBlended == 1 {
				retColor = retColor.BlendRgb(col, 0.333)
			} else if colorsBlended == 2 {
				retColor = retColor.BlendRgb(col, 0.25)
			}
			colorsBlended++
		}
	}
	newRed := retColor.R * 255.0
	newGreen := retColor.G * 255.0
	newBlue := retColor.B * 255.0
	cell := uint32(newRed)<<24 + uint32(newGreen)<<16 + uint32(newBlue)<<8 + (oldCell & ALIVE)
	return cell
}

func (dg DataGrid) NeighborsColorMajority(y, x uint32, neighbors byte) uint32 {
	colorCount := make(map[colorful.Color]int)
	for i := N; i < 9; i++ {
		if neighbors&DirectionMasks[i] > 0 {
			xOff := XOffsets[i]
			yOff := YOffsets[i]
			cell := dg[int(y)+yOff][int(x)+xOff]
			col := colorOfCell(cell)
			colorCount[col]++
		}
	}
	lastColor := colorful.Color{}
	for col, count := range colorCount {
		if count == 2 || count > 2 {
			newRed := col.R * 255.0
			newGreen := col.G * 255.0
			newBlue := col.B * 255.0
			return uint32(newRed)<<24 + uint32(newGreen)<<16 + uint32(newBlue)<<8 + ALIVE
		} else {
			lastColor = col
		}
	}
	//Just pick the last color in the map (maps aren't sorted, so this should be relatively random)
	newRed := lastColor.R * 255.0
	newGreen := lastColor.G * 255.0
	newBlue := lastColor.B * 255.0
	return uint32(newRed)<<24 + uint32(newGreen)<<16 + uint32(newBlue)<<8 + ALIVE
}

func colorOfCell(cell uint32) colorful.Color {
	return colorful.Color{R: float64((cell>>24)&0x000000FF) / 255.0, G: float64((cell>>16)&0x000000FF) / 255.0, B: float64((cell>>8)&0x000000FF) / 255.0}
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

var XOffsets = map[Direction]int{
	NW: -1,
	N:  0,
	NE: 1,
	E:  1,
	SE: 1,
	S:  0,
	SW: -1,
	W:  -1,
}
var YOffsets = map[Direction]int{
	NW: -1,
	N:  -1,
	NE: -1,
	E:  0,
	SE: 1,
	S:  1,
	SW: 1,
	W:  0,
}
var DirectionMasks = map[Direction]byte{
	NW: 0b0000_0001,
	N:  0b0000_0010,
	NE: 0b0000_0100,
	E:  0b0000_1000,
	SE: 0b0001_0000,
	S:  0b0010_0000,
	SW: 0b0100_0000,
	W:  0b1000_0000,
}

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
