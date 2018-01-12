/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

package textdraw

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	Renderer *sdl.Renderer
	Width    int32
	Height   int32
}

func NewWindow(r *sdl.Renderer, w, h int32) *Window {
	return &Window{Renderer: r, Width: w, Height: h}
}

func (w *Window) SetWidth(width int32) {
    w.Width = width
}

func (w *Window) SetHeight(height int32) {
    w.Height = height
}

func (w *Window) GetWidth() int32 {
    return w.Width
}

func (w *Window) GetHeight() int32 {
    return w.Height
}
