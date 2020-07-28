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
	for y := int64(0); y < world.height; y++ {
		for x := int64(0); x < world.width; x++ {
			alive := isAlive((*world.data)[y][x])

			neighborhood := world.data.NeighborsValue(y, x)

			(*world.dataBuffer)[y][x] = ConwayNewState(alive, neighborhood)
		}
	}
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
