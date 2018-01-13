/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

package main

import(
    "github.com/veandco/go-sdl2/sdl"
    "log"
    "fmt"
)

//Main gameloop
func gameloop(r *sdl.Renderer) {
}

//Delete this when done
func gameloopDebug(r *sdl.Renderer) {
	err := renderSplash(r)
    if err != nil {
        log.Fatal(fmt.Errorf("Error: could not render splash screen - %v", err))
    }
	renderMenuScreen(r)
}
