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
    
    //TODO: input actual coordinates
	coords[0] = [2]int32{0, 0}
	coords[1] = [2]int32{0, 0}
	coords[2] = [2]int32{0, 0}
	coords[3] = [2]int32{0, 0}
	coords[4] = [2]int32{0, 0}
	coords[5] = [2]int32{0, 0}
	coords[6] = [2]int32{0, 0}
	coords[7] = [2]int32{0, 0}
	coords[8] = [2]int32{0, 0}
	coords[9] = [2]int32{0, 0}
	coords[10] = [2]int32{0, 0}
	coords[11] = [2]int32{0, 0}
	coords[12] = [2]int32{0, 0}
	coords[13] = [2]int32{0, 0}
	coords[14] = [2]int32{0, 0}
	coords[15] = [2]int32{0, 0}
	coords[16] = [2]int32{0, 0}
	coords[17] = [2]int32{0, 0}
	coords[18] = [2]int32{0, 0}
	coords[19] = [2]int32{0, 0}
	coords[20] = [2]int32{0, 0}
	coords[21] = [2]int32{0, 0}
	coords[22] = [2]int32{0, 0}
	coords[23] = [2]int32{0, 0}
	coords[24] = [2]int32{0, 0}
	coords[25] = [2]int32{0, 0}
	coords[26] = [2]int32{0, 0}
	coords[27] = [2]int32{0, 0}
	coords[28] = [2]int32{0, 0}
	coords[29] = [2]int32{0, 0}
	coords[30] = [2]int32{0, 0}
	coords[31] = [2]int32{0, 0}
	coords[32] = [2]int32{0, 0}
	coords[33] = [2]int32{0, 0}

	err := p.setIntersections(coords)
	return err
}
