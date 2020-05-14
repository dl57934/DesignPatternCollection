package observer

type MakingChange struct {
	son, daughter int
	observers     []Observer
}

func (m *MakingChange) makingSon() {
	m.son++
	m.allNotify()
}

func (m *MakingChange) makingDaughter() {
	m.daughter++
	m.allNotify()
}

func (m *MakingChange) addObserver(observer Observer) {
	m.observers = append(m.observers, observer)
}

func (m *MakingChange) allNotify() {
	for _, observer := range m.observers {
		observer.update(m)
	}
}
