package simulation

import (
	"bytes"
	"fmt"
	"github.com/denverquane/golife/proto/message"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type RLE struct {
	name   string
	width  uint32
	height uint32

	data [][]bool
}

func LoadRLE(path string) (RLE, error) {
	rle := RLE{}
	f, err := os.Open(path)
	if err != nil {
		return rle, err
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return rle, nil
	}

	offset := 0
	lines := strings.Split(string(buf), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "#N") {
			rle.name = strings.Replace(line, "#N ", "", 1)
			rle.name = strings.Replace(rle.name, "#N", "", 1)
		} else if line[0] == 'x' || line[0] == 'X' {
			y, x := parseDimLine(strings.ToLower(line))
			rle.height = y
			rle.width = x
		}
		if rle.width > 0 && rle.height > 0 {
			offset = i + 1
			break
		}
	}
	rle.data = make([][]bool, rle.height)
	for y := 0; y < int(rle.height); y++ {
		rle.data[y] = make([]bool, rle.width)
	}
	x := 0
	y := 0
	for _, line := range lines[offset:] {
		buf := bytes.Buffer{}
		for _, c := range line {
			if c == '$' {
				if buf.Len() > 0 {
					length, err := strconv.Atoi(buf.String())
					if err != nil {
						return rle, err
					}
					for i := 0; i < length; i++ {
						y++
					}
				} else {
					y++
				}
				x = 0
				buf = bytes.Buffer{}
			} else if c == 'b' {
				if buf.Len() > 0 {
					length, err := strconv.Atoi(buf.String())
					if err != nil {
						return rle, err
					}
					for i := 0; i < length; i++ {
						rle.data[y][x] = false
						x++
					}
				} else {
					rle.data[y][x] = false
					x++
				}
				buf = bytes.Buffer{}
			} else if c == 'o' {
				if buf.Len() > 0 {
					length, err := strconv.Atoi(buf.String())
					if err != nil {
						return rle, err
					}
					for i := 0; i < length; i++ {
						rle.data[y][x] = true
						x++
					}

				} else {
					rle.data[y][x] = true
					x++
				}
				buf = bytes.Buffer{}
			} else {
				buf.WriteByte(byte(c))
			}
		}
	}
	return rle, nil
}

func parseDimLine(line string) (y, x uint32) {
	line = strings.ReplaceAll(line, " ", "")
	split := strings.Split(line, ",")
	xx, err := strconv.ParseUint(strings.ReplaceAll(split[0], "x=", ""), 10, 64)
	if err != nil {
		log.Println(err)
		xx = 0
	}
	yy, err := strconv.ParseUint(strings.ReplaceAll(split[1], "y=", ""), 10, 64)
	if err != nil {
		log.Println(err)
		yy = 0
	}
	return uint32(yy), uint32(xx)
}

func (rle RLE) ToString() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("Name: %s, Height: %d, Width: %d\n", rle.name, rle.height, rle.width))
	for y := 0; y < int(rle.height); y++ {
		for x := 0; x < int(rle.width); x++ {
			if rle.data[y][x] {
				buf.WriteString(" X ")
			} else {
				buf.WriteString(" _ ")
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (rle RLE) toProto() *message.RLE {
	//TODO better packing? 8 cells per byte, not 1?
	data := make([]byte, rle.height*rle.width)
	idx := 0
	for _, row := range rle.data {
		for _, cell := range row {
			if cell {
				data[idx] = 0xFF
			} else {
				data[idx] = 0x00
			}
			idx++
		}
	}
	return &message.RLE{
		Name:   rle.name,
		Width:  rle.width,
		Height: rle.height,
		Data:   data,
	}
}

func ToRleBytes(rles map[string]RLE) []byte {
	rleArr := make([]RLE, len(rles))
	idx := 0
	for name, rle := range rles {
		rle.name = name
		rleArr[idx] = rle
		idx++
	}
	msg := message.RLEs{}
	for _, rle := range rleArr {
		msg.Rles = append(msg.Rles, rle.toProto())
	}
	rlesMarshalled, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
	}
	return rlesMarshalled
}
