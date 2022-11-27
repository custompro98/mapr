package grid

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/ui"
)

var help = []string{
	"h: move left",
	"j: move down",
	"k: move up",
	"l: move right",
	"u: undo",
	"?: help",
	"q: quit",
}

var empty = gruid.Cell{Rune: ' '}
var visited = gruid.Cell{Rune: '.'}
var active = gruid.Cell{Rune: 'O'}
var wall = gruid.Cell{Rune: 'X'}

type dirName = int

const (
	N dirName = iota
	E
	S
	W
)

type direction struct {
	name dirName
	x    int
	y    int
}

var north = direction{
	name: N,
	x:    0,
	y:    -1,
}
var east = direction{
	name: E,
	x:    1,
	y:    0,
}
var south = direction{
	name: S,
	x:    0,
	y:    1,
}
var west = direction{
	name: W,
	x:    -1,
	y:    0,
}

type model struct {
	display gruid.Grid
	mapr    *gruid.Grid
	pos     gruid.Point
	bearing direction
	history []model
	help    *ui.Pager
}

func newModel(gd gruid.Grid) *model {
	m := &model{
		display: gd,
		mapr:    &gd,
		pos: gruid.Point{
			X: 0,
			Y: gd.Ug.Height / 2,
		},
		bearing: east,
		help: newPager(PagerConfig{
			w:     gd.Ug.Width,
			h:     gd.Ug.Height,
			title: "Help",
			body:  help,
		}),
	}

	return m
}

func (m *model) initialize() {
	m.move(m.pos)
}

// Draw implements gruid.Model#Draw
func (m *model) Draw() gruid.Grid {
	return m.display
}

// Update implements gruid.Model#Update
func (m *model) Update(msg gruid.Msg) gruid.Effect {
	switch msg := msg.(type) {
	case gruid.MsgInit:
		m.initialize()
	case gruid.MsgKeyDown:
		return m.updateMsgKeyDown(msg)
	}

	return nil
}

func (m *model) updateMsgKeyDown(msg gruid.MsgKeyDown) gruid.Effect {
	m.display = *m.mapr

	switch msg.Key {
	case gruid.KeyArrowLeft, "h", "H":
		m.bearing = west
	case gruid.KeyArrowDown, "j", "J":
		m.bearing = south
	case gruid.KeyArrowUp, "k", "K":
		m.bearing = north
	case gruid.KeyArrowRight, "l", "L":
		m.bearing = east
	case "r", "R":
		m.buildRoom()
		return nil
	case "u", "U":
		m.undo()
		return nil
	case "?":
		m.display = m.help.Draw()
		return nil
	case gruid.KeyEscape, "Q", "q":
		return gruid.End()
	default:
		return nil
	}

	m.move(m.pos.Shift(m.bearing.x, m.bearing.y))

	return nil
}

func (m *model) move(to gruid.Point) {
	if !m.passable(to) {
		return
	}

	from := m.pos

	m.mapr.Set(from, visited)
	m.mapr.Set(to, active)

	m.history = append(m.history, *m)
	m.pos = to

	m.buildWalls(m.pos)
}

func (m *model) passable(p gruid.Point) bool {
	return m.mapr.Contains(p)
}

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

func (m *model) undo() {
	if len(m.history) == 0 {
		return
	}

	last := m.history[len(m.history)-1]

	m.mapr = last.mapr
	m.display = *m.mapr
	m.pos = last.pos
	m.bearing = last.bearing
	m.history = last.history
}

func (m *model) buildRoom() {
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

func (m *model) visit(p gruid.Point) {
	m.mapr.Set(p, visited)
}
