package Observer

type MakingChange struct {
	son, daughter int
	observers     []Observer
}

func (m *MakingChange) MakingSon() {
	m.son++
	m.AllNotify()
}

func (m *MakingChange) MakingDaughter() {
	m.daughter++
	m.AllNotify()
}

func (m *MakingChange) AddObserver(observer Observer) {
	m.observers = append(m.observers, observer)
}

func (m *MakingChange) AllNotify() {
	for _, observer := range m.observers {
		observer.update(m)
	}
}
