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
	"time"
	"textdraw"
)

//Create and display splash screen for three seconds
func renderSplash(r *sdl.Renderer) error {
	r.Clear()

	err := drawLogo(&textdraw.Window{Renderer: r, Width: Width, Height: Height},(Width-287)/2,(Height-300)/2)
	if err != nil {
		return fmt.Errorf("Error: could not render splash screen - %v", err)
	}

	r.Present()

	time.Sleep(time.Millisecond * 3000)

	return nil
}
