package grid

import "github.com/anaseto/gruid"

func (m *model) buildWalls(p gruid.Point) {
	m.buildWall(p.Shift(north.x, north.y))
	m.buildWall(p.Shift(east.x, east.y))
	m.buildWall(p.Shift(south.x, south.y))
	m.buildWall(p.Shift(west.x, west.y))
}

func (m *model) buildWall(p gruid.Point) {
	if m.mapr.At(p).Rune != empty.Rune {
		return
	}

	m.mapr.Set(p, wall)
}
