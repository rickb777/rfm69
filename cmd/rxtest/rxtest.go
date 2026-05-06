package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/rickb777/gpio/sysfs"
	"github.com/rickb777/rfm69"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.LUTC)
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s frequency", os.Args[0])
	}

	sysfs.Debugf = log.Printf // enable debug output

	frequency := getFrequency(os.Args[1])
	r := rfm69.Open()
	must(r.Error())

	log.Printf("setting frequency to %d", frequency)
	r.Init(frequency, 76800, 120_000)
	for r.Error() == nil {
		data, rssi := r.Receive(time.Hour)
		log.Printf("% X (RSSI = %d)", data, rssi)
	}
	must(r.Error())
}

func getFrequency(s string) uint32 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	if 430.0 <= f && f <= 920.0 {
		return uint32(f * 1000000.0)
	}
	if 430000000.0 <= f && f <= 920000000.0 {
		return uint32(f)
	}
	log.Fatalf("%s: invalid pump frequency", s)
	panic("unreachable")
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
