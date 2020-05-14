package observer

import (
	"fmt"
	"strconv"
)

type Younghee struct {
}

func (y Younghee) update(makingChange *MakingChange) {
	fmt.Println("Younghee has son: " + strconv.Itoa(makingChange.son) + " daughter: " + strconv.Itoa(makingChange.daughter))
}
