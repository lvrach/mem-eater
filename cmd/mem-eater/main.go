package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	humanize "github.com/dustin/go-humanize"
)

var rateFlag = flag.String("rate", "2MB", "rate of memory consumsion per second")
var maxFlag = flag.String("max", "500MB", "max memory consumsion")

// prevents garbage collection
var allocated [][]byte

func allocate(rate, max uint64) {
	allocated = make([][]byte, 0)
	for consumedBytes := uint64(0); consumedBytes < max; consumedBytes += rate {

		chunk := make([]byte, rate)
		// trick operating system into allocation this section in physical memory
		for i := uint64(0); i < rate; i++ {
			chunk[i] = byte(1)
		}

		allocated = append(allocated, make([]byte, rate))
		time.Sleep(time.Second)
	}
}

func main() {
	flag.Parse()

	rate, err := humanize.ParseBytes(*rateFlag)
	if err != nil {
		log.Fatalln(err)
	}

	max, err := humanize.ParseBytes(*maxFlag)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("start consuming %s memory per second, until we reach %s\n", *rateFlag, *maxFlag)

	allocate(rate, max)

	log.Printf("max memory consumsion of %s reached, press a key to exit:\n", *maxFlag)
	fmt.Scanln()
}
