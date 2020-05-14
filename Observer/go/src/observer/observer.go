package Observer

type Observer interface {
	update(makingChange *MakingChange)
}
