package simulation

import (
	"bytes"
	"fmt"
	"github.com/denverquane/golife/proto/message"
	"google.golang.org/protobuf/proto"
)

type DataGrid [][]uint32

const ALIVE uint32 = 0x00_00_00_FF
const ALIVE_MASK uint32 = 0xFF_FF_FF_00
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
		data = append(data, (*world.data)[y]...)
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

func Decay(cell uint32) uint32 {
	lowerByte := cell & ALIVE
	//log.Printf("%32b\n", lowerByte)
	if lowerByte == uint32(1) {
		//don't kill the cell, just keep it at 1
		return cell
	} else {
		return cell - 1
	}
}

func (world *World) Tick() {
	//only span the inner areas, don't do the perimeter yet
	for y := uint32(1); y < world.height-1; y++ {
		for x := uint32(1); x < world.width-1; x++ {
			alive := isAliveBool((*world.data)[y][x])
			neighborhood := world.data.InnerNeighborsValue(y, x)

			if alive {
				if world.aliveRulesMapping[neighborhood] {
					(*world.dataBuffer)[y][x] = Decay((*world.data)[y][x])
				} else {
					(*world.dataBuffer)[y][x] = 0
				}
			} else {
				if world.deadRulesMapping[neighborhood] {
					(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorAverage(y, x, neighborhood)
				} else {
					(*world.dataBuffer)[y][x] = DEAD
				}
			}

		}
	}
	y := uint32(0)
	for x := uint32(1); x < world.width-1; x++ {
		alive := isAliveBool((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(N, y, x)
		if alive {
			if world.aliveRulesMapping[neighborhood] {
				(*world.dataBuffer)[y][x] = Decay((*world.data)[y][x])
			} else {
				(*world.dataBuffer)[y][x] = 0
			}
		} else {
			if world.deadRulesMapping[neighborhood] {
				(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorAverage(y, x, neighborhood)
			} else {
				(*world.dataBuffer)[y][x] = DEAD
			}
		}
	}

	y = world.height - 1
	for x := uint32(1); x < world.width-1; x++ {
		alive := isAliveBool((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(S, y, x)
		if alive {
			if world.aliveRulesMapping[neighborhood] {
				(*world.dataBuffer)[y][x] = Decay((*world.data)[y][x])
			} else {
				(*world.dataBuffer)[y][x] = 0
			}
		} else {
			if world.deadRulesMapping[neighborhood] {
				(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorAverage(y, x, neighborhood)
			} else {
				(*world.dataBuffer)[y][x] = DEAD
			}
		}
	}

	x := world.width - 1
	for y := uint32(1); y < world.height-1; y++ {
		alive := isAliveBool((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(E, y, x)
		if alive {
			if world.aliveRulesMapping[neighborhood] {
				(*world.dataBuffer)[y][x] = Decay((*world.data)[y][x])
			} else {
				(*world.dataBuffer)[y][x] = 0
			}
		} else {
			if world.deadRulesMapping[neighborhood] {
				(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorAverage(y, x, neighborhood)
			} else {
				(*world.dataBuffer)[y][x] = DEAD
			}
		}
	}

	x = uint32(0)
	for y := uint32(1); y < world.height-1; y++ {
		alive := isAliveBool((*world.data)[y][x])
		neighborhood := world.data.PerimeterNeighborsValue(W, y, x)
		if alive {
			if world.aliveRulesMapping[neighborhood] {
				(*world.dataBuffer)[y][x] = Decay((*world.data)[y][x])
			} else {
				(*world.dataBuffer)[y][x] = 0
			}
		} else {
			if world.deadRulesMapping[neighborhood] {
				(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorAverage(y, x, neighborhood)
			} else {
				(*world.dataBuffer)[y][x] = DEAD
			}
		}
	}

	//x is already 0
	y = 0
	alive := isAliveBool((*world.data)[y][x])
	neighborhood := world.data.PerimeterNeighborsValue(NW, y, x)
	if alive {
		if world.aliveRulesMapping[neighborhood] {
			(*world.dataBuffer)[y][x] = Decay((*world.data)[y][x])
		} else {
			(*world.dataBuffer)[y][x] = 0
		}
	} else {
		if world.deadRulesMapping[neighborhood] {
			(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorAverage(y, x, neighborhood)
		} else {
			(*world.dataBuffer)[y][x] = DEAD
		}
	}

	x = world.width - 1
	alive = isAliveBool((*world.data)[y][x])
	neighborhood = world.data.PerimeterNeighborsValue(NE, y, x)
	if alive {
		if world.aliveRulesMapping[neighborhood] {
			(*world.dataBuffer)[y][x] = Decay((*world.data)[y][x])
		} else {
			(*world.dataBuffer)[y][x] = 0
		}
	} else {
		if world.deadRulesMapping[neighborhood] {
			(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorAverage(y, x, neighborhood)
		} else {
			(*world.dataBuffer)[y][x] = DEAD
		}
	}

	y = world.height - 1
	alive = isAliveBool((*world.data)[y][x])
	neighborhood = world.data.PerimeterNeighborsValue(SE, y, x)
	if alive {
		if world.aliveRulesMapping[neighborhood] {
			(*world.dataBuffer)[y][x] = Decay((*world.data)[y][x])
		} else {
			(*world.dataBuffer)[y][x] = 0
		}
	} else {
		if world.deadRulesMapping[neighborhood] {
			(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorAverage(y, x, neighborhood)
		} else {
			(*world.dataBuffer)[y][x] = DEAD
		}
	}

	x = 0
	alive = isAliveBool((*world.data)[y][x])
	neighborhood = world.data.PerimeterNeighborsValue(SW, y, x)
	if alive {
		if world.aliveRulesMapping[neighborhood] {
			(*world.dataBuffer)[y][x] = Decay((*world.data)[y][x])
		} else {
			(*world.dataBuffer)[y][x] = 0
		}
	} else {
		if world.deadRulesMapping[neighborhood] {
			(*world.dataBuffer)[y][x] = (*world.data).NeighborsColorAverage(y, x, neighborhood)
		} else {
			(*world.dataBuffer)[y][x] = DEAD
		}
	}

	tempPtr := world.data
	world.data = world.dataBuffer
	world.dataBuffer = tempPtr
	world.tick++
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

func (world *World) MarkAlive(y, x uint32) {
	(*world.data)[y][x] = ALIVE_FULL
}

func (world *World) MarkAliveColor(y, x uint32, color uint32) {
	(*world.data)[y][x] = ALIVE | (color & ALIVE_MASK)
}

func (world *World) MakeGliderGun(y, x uint32) {
	(*world.data)[y+4][x] = ALIVE_FULL
	(*world.data)[y+4][x+1] = ALIVE_FULL
	(*world.data)[y+5][x] = ALIVE_FULL
	(*world.data)[y+5][x+1] = ALIVE_FULL

	(*world.data)[y+2][x+12] = ALIVE_FULL
	(*world.data)[y+2][x+13] = ALIVE_FULL
	(*world.data)[y+3][x+11] = ALIVE_FULL
	(*world.data)[y+3][x+15] = ALIVE_FULL
	(*world.data)[y+4][x+10] = ALIVE_FULL
	(*world.data)[y+4][x+16] = ALIVE_FULL
	(*world.data)[y+5][x+10] = ALIVE_FULL
	(*world.data)[y+5][x+14] = ALIVE_FULL
	(*world.data)[y+5][x+16] = ALIVE_FULL
	(*world.data)[y+5][x+17] = ALIVE_FULL
	(*world.data)[y+6][x+10] = ALIVE_FULL
	(*world.data)[y+6][x+16] = ALIVE_FULL
	(*world.data)[y+7][x+11] = ALIVE_FULL
	(*world.data)[y+7][x+15] = ALIVE_FULL
	(*world.data)[y+8][x+12] = ALIVE_FULL
	(*world.data)[y+8][x+13] = ALIVE_FULL

	(*world.data)[y][x+24] = ALIVE_FULL
	(*world.data)[y+1][x+22] = ALIVE_FULL
	(*world.data)[y+1][x+24] = ALIVE_FULL
	(*world.data)[y+2][x+20] = ALIVE_FULL
	(*world.data)[y+2][x+21] = ALIVE_FULL
	(*world.data)[y+3][x+20] = ALIVE_FULL
	(*world.data)[y+3][x+21] = ALIVE_FULL
	(*world.data)[y+4][x+20] = ALIVE_FULL
	(*world.data)[y+4][x+21] = ALIVE_FULL
	(*world.data)[y+5][x+22] = ALIVE_FULL
	(*world.data)[y+5][x+24] = ALIVE_FULL
	(*world.data)[y+6][x+24] = ALIVE_FULL

	(*world.data)[y+2][x+34] = ALIVE_FULL
	(*world.data)[y+2][x+35] = ALIVE_FULL
	(*world.data)[y+3][x+34] = ALIVE_FULL
	(*world.data)[y+3][x+35] = ALIVE_FULL
}

const (
	TOGGLE_PAUSE int = 1
	MARK_CELL    int = 2
)

type SimulatorMessage struct {
	Type  int
	X     uint32
	Y     uint32
	Color uint32
}
