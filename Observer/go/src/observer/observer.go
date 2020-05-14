package observer

type Observer interface {
	update(makingChange *MakingChange)
}
