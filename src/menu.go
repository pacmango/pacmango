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

var textColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}

func renderMenuScreen(w *textdraw.Window) error {
	menuFont, err := textdraw.NewTTFFont(MenuFontAddress, 20, textColor)
	if err != nil {
		return fmt.Errorf("Error: could not cretae new font - %v", err)
	}
	defer menuFont.CloseTTFFont()
	w.Renderer.Clear()
	w.Renderer.Present()
	return nil
}
