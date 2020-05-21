package Test

import "fmt"

type CharParent struct {
	Letter string
}

func (char CharParent) Open() {
	fmt.Println("open")
}

func (char CharParent) Close() {
	fmt.Println("close")
}

func (char CharParent) Print() {
	fmt.Println(char.Letter)
}
