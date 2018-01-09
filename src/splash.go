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
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func renderSplash(r *sdl.Renderer) error {
	r.Clear()

	r.SetDrawColor(255, 0, 255, 255)
	r1 := &sdl.Rect{X: 100, Y: 100, W: 55, H: 55}
	r.DrawRect(r1)
	r.FillRect(r1)
	r.SetDrawColor(0, 0, 0, 255)

	r.Present()

	time.Sleep(time.Millisecond * 3000)

	return nil
}

func gameloopDebug(r *sdl.Renderer) {
	renderSplash(r)
	r.Clear()
	r.Present()
}
