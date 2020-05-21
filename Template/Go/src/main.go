package main

import "main.go/Test"

func main() {
	stringParent := Test.StringParent{"test"}
	charParent := Test.CharParent{"test"}
	writeLine(stringParent, charParent)
}

func writeLine(parents ...Test.Parent) {

	for _, parent := range parents {
		parent.Open()
		parent.Print()
		parent.Close()
	}
}
