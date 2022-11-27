package grid

import "github.com/anaseto/gruid"

func (m *model) move(to gruid.Point) {
	if !m.passable(to) {
		return
	}

	m.snapshot()

	from := m.pos

	m.mapr.Set(from, visited)
	m.mapr.Set(to, active)

	m.pos = to

	m.buildWalls(m.pos)
}

func (m *model) visit(p gruid.Point) {
	m.mapr.Set(p, visited)
}

func (m *model) passable(p gruid.Point) bool {
	return m.mapr.Contains(p)
}

