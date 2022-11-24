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
	path    []gruid.Point
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

	m.path = append(m.path, from)
	m.pos = to

	m.erectWalls()
}

func (m *model) passable(p gruid.Point) bool {
	return m.mapr.Contains(p)
}

func (m *model) erectWalls() {
	north := m.pos.Shift(north.x, north.y)
	east := m.pos.Shift(east.x, east.y)
	south := m.pos.Shift(south.x, south.y)
	west := m.pos.Shift(west.x, west.y)

	if m.mapr.At(north).Rune == empty.Rune {
		m.mapr.Set(north, wall)
	}

	if m.mapr.At(east).Rune == empty.Rune {
		m.mapr.Set(east, wall)
	}

	if m.mapr.At(south).Rune == empty.Rune {
		m.mapr.Set(south, wall)
	}

	if m.mapr.At(west).Rune == empty.Rune {
		m.mapr.Set(west, wall)
	}
}

func (m *model) undo() {
	last := m.path[len(m.path)-1]

	m.mapr.Set(m.pos, empty)
	m.mapr.Set(last, active)

	m.pos = last

	m.erectWalls()
}
