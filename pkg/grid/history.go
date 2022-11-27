package grid

func (m *model) snapshot() {
	m.history = append(m.history, *m)
}

func (m *model) undo() {
	if len(m.history) == 0 {
		return
	}

	last := m.history[len(m.history)-1]

	prev := empty

	if len(m.history) > 1 {
		prev = m.mapr.At(m.history[len(m.history)-2].pos)
	}

	last.mapr.Set(m.pos, prev)
	last.mapr.Set(last.pos, active)

	m.mapr = last.mapr
	m.display = *m.mapr
	m.pos = last.pos
	m.bearing = last.bearing
	m.history = last.history
}
