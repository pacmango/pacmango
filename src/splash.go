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
	"github.com/veandco/go-sdl2/ttf"
	"time"
	"fmt"
)

var titleTextLines = []string{"Spectre","Chase"}
var titleTextColor = sdl.Color{R: 250, G: 206, B: 52, A: 255}
var titleTextVerticalPos = []int32{100, 200, 300}

//Create and display splash screen for three seconds
func renderSplash(r *sdl.Renderer) error {
	//Load font
	titleFont, err := ttf.OpenFont(TitleFontAddress,90)
	if err != nil {
		return fmt.Errorf("could not open title font")
	}
	titleFont.SetOutline(1)
	defer titleFont.Close()

	//Create title text surfaces
	titleTextSurfaces := make([]*sdl.Surface,len(titleTextLines))

	for index, text := range titleTextLines {
		titleTextSurfaces[index], err = titleFont.RenderUTF8Solid(text, titleTextColor)
		if err != nil {
			return fmt.Errorf("could not render title text surface for line %d", index+1)
		}
	}

	//Create title text textures and render
	r.Clear()

	for index, surface := range titleTextSurfaces {
		texture, err := r.CreateTextureFromSurface(surface)
		if err != nil {
			return fmt.Errorf("could not create title text texture from surface %d", index+1)
		}
		r.Copy(texture,nil,&sdl.Rect{X:50,Y:titleTextVerticalPos[index],W:Width-100,H:100})
	}

	r.Present()

	time.Sleep(time.Millisecond * 3000)

	return nil
}
