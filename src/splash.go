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
	"time"
)

//TODO: Add yellow text with red border instead of hollow yellow text

var titleTextLines = []string{"Spectre", "Chase"}
var titleTextColor = sdl.Color{R: 250, G: 206, B: 52, A: 255}
var titleTextOutlineColor = sdl.Color{R: 175, G: 0, B: 0, A: 255}
var titleTextVerticalPos = []int32{100, 200, 300}

//Create and display splash screen for three seconds
func renderSplash(r *sdl.Renderer) error {
	//Load font
	titleFont, err := ttf.OpenFont(TitleFontAddress, 200)
	if err != nil {
		return fmt.Errorf("could not open title font")
	}
	defer titleFont.Close()

	r.Clear()

	//Create and render title text with border
	for index, text := range titleTextLines {
		titleFont.SetOutline(0)
		textSurface, err := titleFont.RenderUTF8Solid(text,titleTextColor)
		if err != nil {
			return fmt.Errorf("could not render title text surface for line %d", index+1)
		}
		defer textSurface.Free()

		titleFont.SetOutline(3)
		textBorderSurface, err := titleFont.RenderUTF8Solid(text,titleTextOutlineColor)
		if err != nil {
			return fmt.Errorf("could not render title text border surface for line %d", index+1)
		}
		defer textBorderSurface.Free()

		textTexture, err := r.CreateTextureFromSurface(textSurface)
		if err != nil {
			return fmt.Errorf("could not create title text texture from surface %d", index+1)
		}
		defer textTexture.Destroy()

		textBorderTexture, err := r.CreateTextureFromSurface(textBorderSurface)
		if err != nil {
			return fmt.Errorf("could not create title border texture from surface %d", index+1)
		}
		defer textTexture.Destroy()

		if err := r.Copy(textTexture, nil, &sdl.Rect{X: 50, Y: titleTextVerticalPos[index], W: Width - 100, H: 100}); err != nil {
			return fmt.Errorf("could not copy texture %d to render target", index+1)
		}
		if err := r.Copy(textBorderTexture, nil, &sdl.Rect{X: 50-2, Y: titleTextVerticalPos[index]-2, W: Width - 100+2, H: 100+2}); err != nil {
			return fmt.Errorf("could not copy texture %d to render target", index+1)
		}

	}

	r.Present()

	time.Sleep(time.Millisecond * 3000)

	return nil
}
