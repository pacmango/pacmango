/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
)

const Width int32 = 448
const Height int32 = 576
const TitleFontAddress string = "fonts/Ceviche_One/CevicheOne-Regular.ttf"
const MenuFontAddress string = "fonts/PressStart2P/PressStart2P-Regular.ttf"

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
		log.Fatal(fmt.Errorf("Error: SDL2 initialization failed - %v", err))
	}
	defer sdl.Quit()

	err = ttf.Init()
	if err != nil {
		log.Fatal(fmt.Errorf("Error: SDL2_TTF initialization failed - %v", err))
	}
	defer ttf.Quit()

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
