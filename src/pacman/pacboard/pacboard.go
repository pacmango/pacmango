/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

//The pacman/packboard package provides for creating, loading, and modifying
//game board data for pacman-like games
package pacboard

import (
	"fmt"
)

//See "getPacboardItem" for names of classes/subclasses/item numbers
type PacboardItem struct {
	Class    string
	Subclass string
	Item     int8
}
type Pacboard [][]PacboardItem

//Marks PacboardItems indicated in intersections as Subclass "Intersection"
//Necessary for proper activation of ghost "AI", as it only runs when the ghost
//is at an intersection.
func (p *Pacboard) setIntersections(intersections [][2]int32) error {
	for _, coord := range intersections {
		if pacitem := p.GetItem(coord[1], coord[0]); pacitem.Class != "Open" {
			return fmt.Errorf("Error: invalid intersection coordinate \"%v\", board not open", coord)
		} else {
			pacitem.Subclass = "Intersection"
			//fmt.Println(p.GetItem(coord[1], coord[0]))
		}
	}
	return nil
}

//Retrieve PacboardItem at column i and row j
func (p *Pacboard) GetItem(i, j int32) *PacboardItem {
	return &(*p)[j][i]
}

//Modify PacboardItem at column i and row j
func (p *Pacboard) SetItem(i, j int32, value PacboardItem) {
	(*p)[j][i] = value
}

//Retrieve the number of columns in Pacboard
func (p *Pacboard) Width() int32 {
	return int32(len((*p)[0]))
}

//Retrieve the number of rows in Pacboard
func (p *Pacboard) Height() int32 {
	return int32(len((*p)))
}

//Checks to make sure row i and column j is within bounds of p.Width() and p.Height()
func (p *Pacboard) ValidPacboardIndex(i, j int32) bool {
	return !(i > p.Width()-1 || i < 0 || j > p.Height()-1 || j < 0)
}

//Creates a new pacboard with given width and height
func NewPacboard(width, height int32) *Pacboard {
	pacboard := make(Pacboard, height)
	for i := int32(0); i < height; i++ {
		pacboard[i] = make([]PacboardItem, width)
	}
	return &pacboard
}

//Creates a new Pacboard from a template file, with the original Pacman dimensions of
//a width of 28 and a height of 36
func NewPacboardFromFile(filePath string) (*Pacboard, error) {
	pacboard := NewPacboard(28, 36)
	err := pacboard.loadFromFile(filePath)
	return pacboard, err
}
