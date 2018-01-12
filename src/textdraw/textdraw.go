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
    *ttf.Font
    color color.RGBA
}

type Window struct {
     renderer *sdl.Renderer
     width int32
     height int32
 }

func NewTTFFont(fontFile string, fontSize int, color color.RGBA) (*Font, error) {
    font, err := ttf.OpenFont(fontFile, fontSize)
    return &Font{font, color}, err
}
