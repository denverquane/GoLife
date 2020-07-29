package simulation

import (
	"bytes"
	"fmt"
)

type DataGrid [][]byte

const ALIVE byte = 0b1000_0000
const DEAD byte = 0b0000_0000

type World struct {
	width  int64
	height int64
	//indexed [y][x]!
	data              *DataGrid
	dataBuffer        *DataGrid
	aliveRulesMapping map[byte]byte
	deadRulesMapping  map[byte]byte
	tick              int64
}

func (world World) GetDims() (height int64, width int64) {
	return world.height, world.width
}

func (world World) GetFlattenedData() []byte {
	return bytes.Join(*world.data, nil)
}

func NewConwayWorld(width, height int64) World {
	data := make(DataGrid, height)
	for i, _ := range data {
		data[i] = make([]byte, width)
	}
	buffer := make(DataGrid, height)
	for i, _ := range buffer {
		buffer[i] = make([]byte, width)
	}
	alive, dead := GenerateConwayNeighborsRules()
	return World{
		width:             width,
		height:            height,
		data:              &data,
		dataBuffer:        &buffer,
		aliveRulesMapping: alive,
		deadRulesMapping:  dead,
		tick:              0,
	}
}

func (world World) GetTick() int64 {
	return world.tick
}

func (world *World) Tick() {

	//only span the inner areas, don't do the perimeter yet
	for y := int64(1); y < world.height-1; y++ {
		for x := int64(1); x < world.width-1; x++ {
			alive := isAlive((*world.data)[y][x])

			neighborhood := world.data.InnerNeighborsValue(y, x)

			(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)
		}
	}
	y := int64(0)
	for x := int64(1); x < world.width-1; x++ {
		alive := isAlive((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(N, y, x)
		(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)
	}

	y = world.height - 1
	for x := int64(1); x < world.width-1; x++ {
		alive := isAlive((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(S, y, x)
		(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)
	}

	x := world.width - 1
	for y := int64(1); y < world.height-1; y++ {
		alive := isAlive((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(E, y, x)
		(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)
	}

	x = int64(0)
	for y := int64(1); y < world.height-1; y++ {
		alive := isAlive((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(W, y, x)
		(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)
	}

	//x is already 0
	y = 0
	alive := isAlive((*world.data)[y][x])
	neighborhood := world.data.PerimeterNeighborsValue(NW, y, x)
	(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)

	x = world.width - 1
	alive = isAlive((*world.data)[y][x])
	neighborhood = world.data.PerimeterNeighborsValue(NE, y, x)
	(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)

	y = world.height - 1
	alive = isAlive((*world.data)[y][x])
	neighborhood = world.data.PerimeterNeighborsValue(SE, y, x)
	(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)

	x = 0
	alive = isAlive((*world.data)[y][x])
	neighborhood = world.data.PerimeterNeighborsValue(SW, y, x)
	(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)

	tempPtr := world.data
	world.data = world.dataBuffer
	world.dataBuffer = tempPtr
	world.tick++
}

func (world World) ToString() string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(fmt.Sprintf("Height: %d, Width: %d\n", world.height, world.width))
	for y := int64(0); y < world.height; y++ {
		for x := int64(0); x < world.width; x++ {
			buf.WriteByte(' ')
			if (*world.data)[y][x]&ALIVE == ALIVE {
				buf.WriteByte('X')
			} else {
				buf.WriteByte('_')
			}
		}
		buf.WriteByte('\n')
	}
	buf.WriteByte('\n')
	return buf.String()
}

func (world *World) MarkAlive(y, x int64) {
	(*world.data)[y][x] = ALIVE
}

func (world *World) MakeGliderGun(y, x int64) {
	(*world.data)[y+4][x] = ALIVE
	(*world.data)[y+4][x+1] = ALIVE
	(*world.data)[y+5][x] = ALIVE
	(*world.data)[y+5][x+1] = ALIVE

	(*world.data)[y+2][x+12] = ALIVE
	(*world.data)[y+2][x+13] = ALIVE
	(*world.data)[y+3][x+11] = ALIVE
	(*world.data)[y+3][x+15] = ALIVE
	(*world.data)[y+4][x+10] = ALIVE
	(*world.data)[y+4][x+16] = ALIVE
	(*world.data)[y+5][x+10] = ALIVE
	(*world.data)[y+5][x+14] = ALIVE
	(*world.data)[y+5][x+16] = ALIVE
	(*world.data)[y+5][x+17] = ALIVE
	(*world.data)[y+6][x+10] = ALIVE
	(*world.data)[y+6][x+16] = ALIVE
	(*world.data)[y+7][x+11] = ALIVE
	(*world.data)[y+7][x+15] = ALIVE
	(*world.data)[y+8][x+12] = ALIVE
	(*world.data)[y+8][x+13] = ALIVE

	(*world.data)[y][x+24] = ALIVE
	(*world.data)[y+1][x+22] = ALIVE
	(*world.data)[y+1][x+24] = ALIVE
	(*world.data)[y+2][x+20] = ALIVE
	(*world.data)[y+2][x+21] = ALIVE
	(*world.data)[y+3][x+20] = ALIVE
	(*world.data)[y+3][x+21] = ALIVE
	(*world.data)[y+4][x+20] = ALIVE
	(*world.data)[y+4][x+21] = ALIVE
	(*world.data)[y+5][x+22] = ALIVE
	(*world.data)[y+5][x+24] = ALIVE
	(*world.data)[y+6][x+24] = ALIVE

	(*world.data)[y+2][x+34] = ALIVE
	(*world.data)[y+2][x+35] = ALIVE
	(*world.data)[y+3][x+34] = ALIVE
	(*world.data)[y+3][x+35] = ALIVE
}
