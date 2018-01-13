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

func renderMenuScreen(w *textdraw.Window) error {
	menuFont, err := textdraw.NewTTFFont(MenuFontAddress, 20, textColor)
	if err != nil {
		return fmt.Errorf("Error: could not cretae new font - %v", err)
	}
	defer menuFont.CloseTTFFont()
	w.Renderer.Clear()
    w.DrawText("Hello",menuFont,20,20)
    w.DrawText("This is the main menu",menuFont,20,50)
    menuFont.SetColor(color.RGBA{R: 240, G: 255, B: 98, A: 255})
    w.DrawText("> Start Game",menuFont,20,80)
	w.Renderer.Present()
	return nil
}
