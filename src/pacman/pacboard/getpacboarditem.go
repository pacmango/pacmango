/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

package pacboard

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
