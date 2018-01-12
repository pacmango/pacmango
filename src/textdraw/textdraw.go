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

func NewWindow(renderer *sdl.Renderer, width, height int32) *Window {
	return &Window{Renderer: renderer, Width: width, Height: height}
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

//Renders text with given font in window w at position (x,y) [Upper-left corner of text]
//Does not call Clear() or Present() on w.Renderer
func (w *Window) DrawText(text string, font *Font, x, y int32) error {
    return nil
}
