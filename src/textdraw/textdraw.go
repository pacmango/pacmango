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
     "github.com/veandco/go-sdl2/ttf"
     "image/color"
 )

type Font struct {
    Font *ttf.Font
    Color color.RGBA
}

type Window struct {
     Renderer *sdl.Renderer
     Width int32
     Height int32
 }

func NewTTFFont(fontFile string, fontSize int, color color.RGBA) (*Font, error) {
    f, err := ttf.OpenFont(fontFile, fontSize)
    return &Font{Font : f, Color : color}, err
}

func (f *Font) CloseTTFFont() {
    f.Font.Close()
    *f = Font{}
}

func (f *Font) SetColor(color color.RGBA) {
    f.Color = color
}

func (f *Font) GetColor() color.RGBA {
    return f.Color
}
