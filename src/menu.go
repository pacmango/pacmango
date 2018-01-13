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
	//"github.com/veandco/go-sdl2/sdl"
	"image/color"
	"textdraw"
)

var textColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

func drawLogo(w *textdraw.Window, x, y int32) error {
	logoTextLines := []string{"Spectre", "Chase"}
	logoTextColor := color.RGBA{R: 250, G: 206, B: 52, A: 255}
	logoTextOutlineColor := color.RGBA{R: 175, G: 0, B: 0, A: 255}
	//logoTextVerticalPos := []int32{100, 200}

	logoFont, err := textdraw.NewTTFFont(TitleFontAddress, 100, logoTextColor)
	if err != nil {
		return fmt.Errorf("Error: could not cretae new font - %v", err)
	}
	defer logoFont.CloseTTFFont()

	width, height, _ := logoFont.Font.SizeUTF8(logoTextLines[0])
	for i, text := range logoTextLines {
		logoFont.Font.SetOutline(0)
		logoFont.Color = logoTextColor
		w.DrawSizedText(text,logoFont, x, y + int32(i)*80, int32(width), int32(height))
		logoFont.Color = logoTextOutlineColor
		logoFont.Font.SetOutline(2)
		w.DrawSizedText(text,logoFont, x - 4, y - 4 + int32(i)*80, int32(width) + 4, int32(height) + 4)
	}

	return nil
}

func renderMenuScreen(w *textdraw.Window) error {
	menuFont, err := textdraw.NewTTFFont(MenuFontAddress, 20, textColor)
	if err != nil {
		return fmt.Errorf("Error: could not cretae new font - %v", err)
	}
	defer menuFont.CloseTTFFont()
	w.Renderer.Clear()
	drawLogo(w,(Width-287)/2,-10)
    w.DrawText("Hello",menuFont,20,170)
    w.DrawText("This is the main menu",menuFont,20,200)
    menuFont.SetColor(color.RGBA{R: 240, G: 255, B: 98, A: 255})
    w.DrawText("> Start Game",menuFont,20,230)
	w.Renderer.Present()
	return nil
}
