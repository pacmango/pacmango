/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

 package main

 import(
     "github.com/veandco/go-sdl2/sdl"
 )

func renderBoard(r *sdl.Renderer) error {
    r.Clear()
    r.Present()
    return nil
}
