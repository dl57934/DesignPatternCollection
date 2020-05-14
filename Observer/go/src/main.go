package main

import "main.go/Observer"

func main() {
	cs := Observer.Chulsu{}
	yh := Observer.Younghee{}
	var mc *Observer.MakingChange = new(Observer.MakingChange)
	mc.AddObserver(cs)
	mc.AddObserver(yh)
	mc.MakingDaughter()
	mc.MakingSon()
}
