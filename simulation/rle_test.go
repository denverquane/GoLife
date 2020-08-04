package simulation

import (
	"log"
	"testing"
)

func TestLoadRLE(t *testing.T) {
	rle, err := LoadRLE("../data/glider.rle")
	if err != nil {
		log.Println(err)
	}
	log.Print(rle.ToString())
}
