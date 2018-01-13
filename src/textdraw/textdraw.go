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
	width, height, err := font.Font.SizeUTF8(text)
	if err != nil {
		return err
	}

	textSurface, err := font.Font.RenderUTF8Solid(text, font.GetSDLColor())
	if err != nil {
		return err
	}
	defer textSurface.Free()

	textTexture, err := w.Renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		return err
	}
	defer textTexture.Destroy()

	if err := w.Renderer.Copy(textTexture, nil, &sdl.Rect{X: x, Y: y, W: int32(width), H: int32(height)}); err != nil {
		return err
	}
    return nil
}

//DrawText() with specified width and height -- will stretch text
func (w *Window) DrawSizedText(text string, font *Font, x, y, width, height int32) error {
	textSurface, err := font.Font.RenderUTF8Solid(text, font.GetSDLColor())
	if err != nil {
		return err
	}
	defer textSurface.Free()

	textTexture, err := w.Renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		return err
	}
	defer textTexture.Destroy()

	if err := w.Renderer.Copy(textTexture, nil, &sdl.Rect{X: x, Y: y, W: width, H: height}); err != nil {
		return err
	}
    return nil
}
