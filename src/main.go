package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

var Width int32 = 448
var Height int32 = 576

func checkQuitEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		return true
	}
	return false
}

func main() {
	//Initialize
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(fmt.Errorf("Error: SDL initialization failed - %v", err))
	}
	defer sdl.Quit()

	//Create window and renderer
	window, err := sdl.CreateWindow("Pacman", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, Width, Height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(fmt.Errorf("Error: Could not create window - %v", err))
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		log.Fatal(fmt.Errorf("Error: Could not create renderer - %v", err))
	}
	defer renderer.Destroy()

	//Run game loop
	go gameloopDebug(renderer)

	//Run window loop
	windowRunning := true
	for event := sdl.PollEvent(); windowRunning; event = sdl.PollEvent() {
		if checkQuitEvent(event) {
			windowRunning = false
		}
	}
}