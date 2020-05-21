package Test

import "fmt"

type StringParent struct {
	Words string
}

func (word StringParent) Open() {
	fmt.Println("Open String")
}

func (word StringParent) Close() {
	fmt.Println("Close String")
}

func (word StringParent) Print() {
	fmt.Println(word.Words)
}
