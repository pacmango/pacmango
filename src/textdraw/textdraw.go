/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

 package textdraw

 import(
     "github.com/veandco/go-sdl2/sdl"
 )

type Window struct {
     Renderer *sdl.Renderer
     Width int32
     Height int32
 }

func NewWindow(r *sdl.Renderer, w, h int32) *Window {
    return &Window{Renderer : r, Width : w, Height : h}
}
