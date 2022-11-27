package grid

import "github.com/anaseto/gruid"

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

