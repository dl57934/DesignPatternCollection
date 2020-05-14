package observer

import (
	"fmt"
	"strconv"
)

type Chulsu struct {
}

func (c Chulsu) update(makingChange *MakingChange) {
	fmt.Println("Chulsu has son: " + strconv.Itoa(makingChange.son) + " daughter: " + strconv.Itoa(makingChange.daughter))
}

func test(){
	fmt.Println("test")
}
