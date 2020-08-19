package simulation

import (
	"bytes"
	"fmt"
	"github.com/denverquane/golife/proto/message"
	"google.golang.org/protobuf/proto"
)

type DataGrid [][]uint32

const ALIVE uint32 = 0x00_00_00_01
const ALIVE_MASK uint32 = 0xFF_FF_FF_FE
const ALIVE_FULL uint32 = 0xFF_FF_FF_FF
const DEAD uint32 = 0x00_00_00_00

type World struct {
	width  uint32
	height uint32
	//indexed [y][x]!
	data              *DataGrid
	dataBuffer        *DataGrid
	aliveRulesMapping map[byte]bool
	deadRulesMapping  map[byte]bool
	tick              uint64
}

func (world *World) GetDims() (height uint32, width uint32) {
	return world.height, world.width
}

func (world *World) GetFlattenedData() []uint32 {
	data := make([]uint32, 0)
	for y := uint32(0); y < world.height; y++ {
		deadCount := uint32(0)
		for x := uint32(0); x < world.width; x++ {
			//if the cell is dead, we can use all the color bits for RLE encoding of sequential dead cells
			if deadCount == world.width-1 || x == world.width-1 {
				shifted := (deadCount << 1) & 0xFFFFFF_FE

				data = append(data, shifted)
				deadCount = 0
			}

			if (*world.data)[y][x]&ALIVE == 0 {
				//dead cell; start the count
				deadCount++
			} else {
				if deadCount > 0 {
					//if we reach an alive cell, add a RLE-encoded length of dead cells
					shifted := DEAD | ((deadCount & 0x000000FF) << 1)
					data = append(data, shifted)

					//append this current alive cell
					data = append(data, (*world.data)[y][x])
					deadCount = 0
				} else {
					//append the alive data normally
					data = append(data, (*world.data)[y][x])
				}
			}
		}
	}
	return data
}

func (world *World) ToMinProtoBytes(paused bool) ([]byte, error) {
	worldMsg := message.WorldData{
		Data:   world.GetFlattenedData(),
		Tick:   world.GetTick(),
		Paused: paused,
	}
	worldMsgMarshalled, err := proto.Marshal(&worldMsg)
	if err != nil {
		return nil, err
	}
	msg := message.Message{
		Type:    message.MessageType_WORLD_DATA,
		Content: worldMsgMarshalled,
	}
	marshalled, err := proto.Marshal(&msg)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}

func (world *World) ToFullProtoBytes() ([]byte, error) {
	worldMsg := message.WorldData{
		Data:   world.GetFlattenedData(),
		Tick:   world.tick,
		Width:  world.width,
		Height: world.height,
	}
	worldMsgMarshalled, err := proto.Marshal(&worldMsg)
	if err != nil {
		return nil, err
	}
	msg := message.Message{
		Type:    message.MessageType_WORLD_DATA,
		Content: worldMsgMarshalled,
	}
	marshalled, err := proto.Marshal(&msg)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}
func NewConwayWorld(width, height uint32) World {
	data := make(DataGrid, height)
	for i, _ := range data {
		data[i] = make([]uint32, width)
	}
	buffer := make(DataGrid, height)
	for i, _ := range buffer {
		buffer[i] = make([]uint32, width)
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

func (world *World) GetTick() uint64 {
	return world.tick
}

func (world *World) Tick(blendColors bool) {
	//only span the inner areas, don't do the perimeter yet
	for y := uint32(1); y < world.height-1; y++ {
		for x := uint32(1); x < world.width-1; x++ {
			alive := isAliveBool((*world.data)[y][x])
			neighborhood := world.data.InnerNeighborsValue(y, x)

			if alive {
				world.setNewAliveBufferState(y, x, neighborhood, blendColors)
			} else {
				world.setNewDeadBufferState(y, x, neighborhood, blendColors)
			}
		}
	}
	y := uint32(0)
	for x := uint32(1); x < world.width-1; x++ {
		alive := isAliveBool((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(N, y, x)
		if alive {
			world.setNewAliveBufferState(y, x, neighborhood, blendColors)
		} else {
			world.setNewDeadBufferState(y, x, neighborhood, blendColors)
		}
	}

	y = world.height - 1
	for x := uint32(1); x < world.width-1; x++ {
		alive := isAliveBool((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(S, y, x)
		if alive {
			world.setNewAliveBufferState(y, x, neighborhood, blendColors)
		} else {
			world.setNewDeadBufferState(y, x, neighborhood, blendColors)
		}
	}

	x := world.width - 1
	for y := uint32(1); y < world.height-1; y++ {
		alive := isAliveBool((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(E, y, x)
		if alive {
			world.setNewAliveBufferState(y, x, neighborhood, blendColors)
		} else {
			world.setNewDeadBufferState(y, x, neighborhood, blendColors)
		}
	}

	x = uint32(0)
	for y := uint32(1); y < world.height-1; y++ {
		alive := isAliveBool((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(W, y, x)
		if alive {
			world.setNewAliveBufferState(y, x, neighborhood, blendColors)
		} else {
			world.setNewDeadBufferState(y, x, neighborhood, blendColors)
		}
	}

	//x is already 0
	y = 0
	alive := isAliveBool((*world.data)[y][x])
	neighborhood := world.data.PerimeterNeighborsValue(NW, y, x)
	if alive {
		world.setNewAliveBufferState(y, x, neighborhood, blendColors)
	} else {
		world.setNewDeadBufferState(y, x, neighborhood, blendColors)
	}

	x = world.width - 1
	alive = isAliveBool((*world.data)[y][x])
	neighborhood = world.data.PerimeterNeighborsValue(NE, y, x)
	if alive {
		world.setNewAliveBufferState(y, x, neighborhood, blendColors)
	} else {
		world.setNewDeadBufferState(y, x, neighborhood, blendColors)
	}

	y = world.height - 1
	alive = isAliveBool((*world.data)[y][x])
	neighborhood = world.data.PerimeterNeighborsValue(SE, y, x)
	if alive {
		world.setNewAliveBufferState(y, x, neighborhood, blendColors)
	} else {
		world.setNewDeadBufferState(y, x, neighborhood, blendColors)
	}

	x = 0
	alive = isAliveBool((*world.data)[y][x])
	neighborhood = world.data.PerimeterNeighborsValue(SW, y, x)
	if alive {
		world.setNewAliveBufferState(y, x, neighborhood, blendColors)
	} else {
		world.setNewDeadBufferState(y, x, neighborhood, blendColors)
	}

	tempPtr := world.data
	world.data = world.dataBuffer
	world.dataBuffer = tempPtr
	world.tick++
}

func (world *World) setNewAliveBufferState(y, x uint32, neighborhood byte, blendColors bool) {
	if world.aliveRulesMapping[neighborhood] {
		if blendColors {
			(*world.dataBuffer)[y][x] = (*world.data).ExistingCellNeighborsColorBlend((*world.data)[y][x], y, x, neighborhood)
		} else {
			(*world.dataBuffer)[y][x] = (*world.data)[y][x]
		}
	} else {
		(*world.dataBuffer)[y][x] = 0
	}
}

func (world *World) setNewDeadBufferState(y, x uint32, neighborhood byte, blendColors bool) {
	if world.deadRulesMapping[neighborhood] {
		if blendColors {
			(*world.dataBuffer)[y][x] = (*world.data).NewCellNeighborsColorBlend(y, x, neighborhood)
		} else {
			(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorMajority(y, x, neighborhood)
		}
	} else {
		(*world.dataBuffer)[y][x] = DEAD
	}
}

func (world *World) ToString() string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(fmt.Sprintf("Height: %d, Width: %d\n", world.height, world.width))
	for y := uint32(0); y < world.height; y++ {
		for x := uint32(0); x < world.width; x++ {
			buf.WriteByte(' ')
			if (*world.data)[y][x]&ALIVE > 0 {
				buf.WriteString(fmt.Sprintf("%3d", (*world.data)[y][x]&ALIVE))
			} else {
				buf.WriteString("___")
			}
		}
		buf.WriteByte('\n')
	}
	buf.WriteByte('\n')
	return buf.String()
}

func (world *World) PlaceRLEAtCoords(rle RLE, y, x, color uint32) bool {
	if y+rle.height > world.height || x+rle.width > world.width {
		return false
	}

	for yy := uint32(0); yy < rle.height; yy++ {
		for xx := uint32(0); xx < rle.width; xx++ {
			if rle.data[yy][xx] {
				(*world.data)[y+yy][x+xx] = color | ALIVE
			}
		}
	}
	return true
}

func (world *World) MarkAlive(y, x uint32) {
	(*world.data)[y][x] = ALIVE_FULL
}

func (world *World) MarkAliveColor(y, x uint32, color uint32) {
	(*world.data)[y][x] = ALIVE | (color & ALIVE_MASK)
}

func (world *World) Clear() {
	for y := uint32(0); y < world.height; y++ {
		for x := uint32(0); x < world.width; x++ {
			(*world.data)[y][x] = DEAD
		}
	}
}

const (
	TOGGLE_PAUSE int = 1
	MARK_CELL    int = 2
	PLACE_RLE    int = 3
	CLEAR_BOARD  int = 4
)

type SimulatorMessage struct {
	Type  int
	X     uint32
	Y     uint32
	Color uint32

	Info string
}
