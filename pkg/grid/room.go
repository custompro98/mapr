package grid

import "github.com/anaseto/gruid"

func (m *model) buildRoom() {
	m.snapshot()

	m.visit(m.pos.Shift(m.bearing.x, m.bearing.y))

	for _, coord := range m.findRoomCoordinates() {
		m.visit(coord)
		m.buildWalls(coord)
	}
}

func (m *model) findRoomCoordinates() []gruid.Point {
	var nw gruid.Point

	switch m.bearing.name {
	case N:
		nw = m.pos.Shift(north.x, north.y).Shift(north.x, north.y).Shift(north.x, north.y).Shift(west.x, west.y)
	case E:
		nw = m.pos.Shift(north.x, north.y).Shift(east.x, east.y)
	case S:
		nw = m.pos.Shift(south.x, south.y).Shift(west.x, west.y)
	case W:
		nw = m.pos.Shift(west.x, west.y).Shift(west.x, west.y).Shift(west.x, west.y).Shift(north.x, north.y)
	default:
		nw = m.pos
	}

	nc := nw.Shift(east.x, east.y)
	ne := nc.Shift(east.x, east.y)

	cw := nw.Shift(south.x, south.y)
	cc := cw.Shift(east.x, east.y)
	ce := cc.Shift(east.x, east.y)

	sw := cw.Shift(south.x, south.y)
	sc := sw.Shift(east.x, east.y)
	se := sc.Shift(east.x, east.y)

	ret := make([]gruid.Point, 0)

	for _, coord := range []gruid.Point{nw, nc, ne, cw, cc, ce, sw, sc, se} {
		if m.mapr.At(coord).Rune == empty.Rune {
			ret = append(ret, coord)
		}
	}

	return ret
}
