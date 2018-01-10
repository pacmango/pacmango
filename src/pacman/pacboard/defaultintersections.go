/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

package pacboard

//Sets intersection subclasses to match those of original Pacman board layout
func (p *Pacboard) SetDefaultIntersections() error {
	coords := make([][2]int32, 34)

	//TODO: reverse coordinates to be col, row like p.GetItem(i,j)
	coords[0]  = [2]int32{4, 6}
	coords[1]  = [2]int32{4, 21}
	coords[2]  = [2]int32{8, 1}
	coords[3]  = [2]int32{8, 6}
	coords[4]  = [2]int32{8, 9}
	coords[5]  = [2]int32{8, 12}
	coords[6]  = [2]int32{8, 15}
	coords[7]  = [2]int32{8, 18}
	coords[8]  = [2]int32{8, 21}
	coords[9]  = [2]int32{8, 26}
	coords[10] = [2]int32{11, 6}
	coords[11] = [2]int32{11, 21}
	coords[12] = [2]int32{14, 12}
	coords[13] = [2]int32{14, 15}
	coords[14] = [2]int32{17, 6}
	coords[15] = [2]int32{17, 9}
	coords[16] = [2]int32{17, 18}
	coords[17] = [2]int32{17, 21}
	coords[18] = [2]int32{20, 9}
	coords[19] = [2]int32{20, 18}
	coords[20] = [2]int32{23, 6}
	coords[21] = [2]int32{23, 9}
	coords[22] = [2]int32{23, 18}
	coords[23] = [2]int32{23, 21}
	coords[24] = [2]int32{26, 6}
	coords[25] = [2]int32{26, 9}
	coords[26] = [2]int32{26, 12}
	coords[27] = [2]int32{26, 15}
	coords[28] = [2]int32{26, 18}
	coords[29] = [2]int32{26, 21}
	coords[30] = [2]int32{29, 3}
	coords[31] = [2]int32{29, 24}
	coords[32] = [2]int32{32, 15}
	coords[33] = [2]int32{32, 18}

	err := p.setIntersections(coords)
	return err
}
