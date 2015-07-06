package main

import (
	"io/ioutil"
	"log"
	"runtime"
	"strings"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	initGraphics()
	loadRom("ibm")
}

func loadRom(romName string) []byte {
	romPath := []string{"./games/", romName}
	rom, err := ioutil.ReadFile(strings.Join(romPath, ""))
	if err != nil {
		log.Fatal(err)
	}

	return rom
}
