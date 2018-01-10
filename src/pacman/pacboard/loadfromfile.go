/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 * This file contains code modified from the Go
 * 1.9.2 standard library package io/ioutil
 * Copyright (c) 2009 The Go Authors.
 * All Rights Reserved.
 * See GO_LICENSE.txt for the Go 1.9.2 license
 **********************************************
 */

package pacboard

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//Loads a Pacboard template file into Pacboard p
func (p *Pacboard) loadFromFile(filePath string) error {
	//Load file
	if ext := filepath.Ext(filePath); ext != ".txt" {
		return fmt.Errorf("Error: cannot read file \"%s\", file must be of type \".txt\"", filePath)
	}

	//Open file and load contents into bytes buffer
	//Based off of ioutil.ReadFile(), ioutil.readAll() from Go 1.9.2
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error: could not open file \"%s\", file may not exist", filePath)
	}
	defer file.Close()

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

	//Write file contents into Pacboard p as PacboardItems
	var row, col int32

	for b, err := buffer.ReadByte(); err != io.EOF; b, err = buffer.ReadByte() {
		if b == 10 { //If character is New Line (ASCII 10)
			row++
			col = 0
		} else {
			p.SetItem(col, row, getPacboardItem(b))
			col++
		}
	}
	return nil
}
