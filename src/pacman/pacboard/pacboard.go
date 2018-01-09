package pacboard

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//Classes: Open; Barrier; Gate
//Subclass of Open: Intersection
//Subclass of Barrier: (Double_/Square_/--)Corner_[A,B,C,D]; Double_[VA,VB,HA,HB]; Horizontal; Vertical; Gate_End_[A,B]
//Items: Empty (-1), Dot (0), Power (1), Other (Else)
type PacboardItem struct {
	Class    string
	Subclass string
	Item     int8
}
type Pacboard [][]PacboardItem

//ASCII symbol to PacboardItem
func getPacboardItem(b byte) PacboardItem {
	switch string(b) {
	case "#":
		return PacboardItem{Class: "Open",
			Subclass: "",
			Item:     -1}

	case "*":
		return PacboardItem{Class: "Open",
			Subclass: "",
			Item:     0}
	case "0":
		return PacboardItem{Class: "Open",
			Subclass: "",
			Item:     1}
	case "W":
		return PacboardItem{Class: "Barrier",
			Subclass: "Double_Corner_A",
			Item:     -1}
	case "X":
		return PacboardItem{Class: "Barrier",
			Subclass: "Double_Corner_B",
			Item:     -1}
	case "Y":
		return PacboardItem{Class: "Barrier",
			Subclass: "Double_Corner_C",
			Item:     -1}
	case "Z":
		return PacboardItem{Class: "Barrier",
			Subclass: "Double_Corner_D",
			Item:     -1}
	case "w":
		return PacboardItem{Class: "Barrier",
			Subclass: "Corner_A",
			Item:     -1}
	case "x":
		return PacboardItem{Class: "Barrier",
			Subclass: "Corner_B",
			Item:     -1}
	case "y":
		return PacboardItem{Class: "Barrier",
			Subclass: "Corner_C",
			Item:     -1}
	case "z":
		return PacboardItem{Class: "Barrier",
			Subclass: "Corner_D",
			Item:     -1}
	case "L":
		return PacboardItem{Class: "Barrier",
			Subclass: "Double_VA",
			Item:     -1}
	case "R":
		return PacboardItem{Class: "Barrier",
			Subclass: "Double_VB",
			Item:     -1}
	case "T":
		return PacboardItem{Class: "Barrier",
			Subclass: "Double_HA",
			Item:     -1}
	case "B":
		return PacboardItem{Class: "Barrier",
			Subclass: "Double_HB",
			Item:     -1}
	case "v":
		return PacboardItem{Class: "Barrier",
			Subclass: "Vertical",
			Item:     -1}
	case "h":
		return PacboardItem{Class: "Barrier",
			Subclass: "Horizontal",
			Item:     -1}
	case "C":
		return PacboardItem{Class: "Barrier",
			Subclass: "Square_Corner_A",
			Item:     -1}
	case "Q":
		return PacboardItem{Class: "Barrier",
			Subclass: "Square_Corner_B",
			Item:     -1}
	case "K":
		return PacboardItem{Class: "Barrier",
			Subclass: "Square_Corner_C",
			Item:     -1}
	case "G":
		return PacboardItem{Class: "Barrier",
			Subclass: "Square_Corner_D",
			Item:     -1}
	case "b":
		return PacboardItem{Class: "Gate",
			Subclass: "",
			Item:     -1}
	case "E":
		return PacboardItem{Class: "Barrier",
			Subclass: "Gate_End_A",
			Item:     -1}
	case "I":
		return PacboardItem{Class: "Barrier",
			Subclass: "Gate_End_B",
			Item:     -1}
	default:
		return PacboardItem{Class: "",
			Subclass: "",
			Item:     -1}
	}
}

//Marks PacboardItems indicated in intersections as Subclass "Intersection"
//Necessary for proper activation of ghost "AI", as it only runs when the ghost
//is at an intersection.
func (p *Pacboard) setIntersections(intersections [][2]int32) error {
	for _, coord := range intersections {
		if pacitem := p.GetItem(coord[0], coord[1]); pacitem.Class != "Open" {
			return fmt.Errorf("Error: invalid intersection coordinate \"%v\", board not open", coord)
		} else {
			pacitem.Subclass = "Intersection"
			fmt.Println(p.GetItem(coord[0], coord[1]))
		}
	}
	return nil
}

//Sets intersection subclasses to match those of original Pacman board layout
func (p *Pacboard) SetDefaultIntersections() error {
	coords := make([][2]int32, 34)
	coords[0] = [2]int32{0, 0}
	err := p.setIntersections(coords)
	return err
}

//Loads a Pacboard template file into Pacboard p
//Based off of ioutil.ReadFile(), ioutil.readAll() from Go 1.9.2
//[Copyright (c) 2009, The Go Authors. All Rights Reserved.]
//[See GO_LICENSE.txt for the BSD license for Go 1.9.2]
func (p *Pacboard) loadFromFile(filePath string) error {
	//Load file
	if ext := filepath.Ext(filePath); ext != ".txt" {
		return fmt.Errorf("Error: cannot read file \"%s\", file must be of type \".txt\"", filePath)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error: could not open file \"%s\", file may not exist", filePath)
	}
	defer file.Close()

	//Read file
	var filesize int64

	if filestat, err := file.Stat(); err != nil {
		return fmt.Errorf("Error: could not get file information")
	} else {
		filesize = filestat.Size()
	}
	buffer := bytes.NewBuffer(make([]byte, 0, filesize))
	if _, err = buffer.ReadFrom(file); err != nil {
		return fmt.Errorf("Error: could not read from file")
	}

	//Write file contents into Pacboard p
	var row, col int32

	for b, err := buffer.ReadByte(); err != io.EOF; b, err = buffer.ReadByte() {
		//If character is New Line (ASCII 10), go to next column
		if b == 10 {
			row++
			col = 0
		} else {
			*(p.GetItem(col, row)) = getPacboardItem(b) //TODO: replace with p.SetItem()
			col++
		}
	}
	return nil
}

//Retrieve PacboardItem at row i and column j
func (p *Pacboard) GetItem(i, j int32) *PacboardItem {
	return &(*p)[j][i]
}

//Modify PacboardItem at row i and column j
func (p *Pacboard) SetItem(i, j int32, value PacboardItem) {

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