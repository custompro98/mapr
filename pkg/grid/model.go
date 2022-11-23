package grid

import (
	"github.com/anaseto/gruid"
)

var empty = gruid.Cell{Rune: ' '}
var visited = gruid.Cell{Rune: '.'}
var active = gruid.Cell{Rune: 'O'}
var wall = gruid.Cell{Rune: 'X'}

type model struct {
  grid gruid.Grid
  pos gruid.Point
  path []gruid.Point
}

func newModel(gd gruid.Grid) *model {
	m := &model{
		grid: gd,
    pos: gruid.Point{
      X: 0,
      Y: gd.Ug.Height / 2,
    },
	}

	return m
}

func (m *model) initialize() {
  m.move(m.pos)
}

// Draw implements gruid.Model#Draw
func (m *model) Draw() gruid.Grid {
	return m.grid
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
  var delta = m.pos

	switch msg.Key {
	case gruid.KeyArrowLeft, "h", "H":
		delta = delta.Shift(-1, 0)
	case gruid.KeyArrowDown, "j", "J":
		delta = delta.Shift(0, 1)
	case gruid.KeyArrowUp, "k", "K":
		delta = delta.Shift(0, -1)
	case gruid.KeyArrowRight, "l", "L":
		delta = delta.Shift(1, 0)
  case "u", "U":
    m.undo(delta)
    return nil
	case gruid.KeyEscape, "Q", "q":
		return gruid.End()
  default:
    return nil
	}

  m.move(delta)

	return nil
}

func (m *model) move(to gruid.Point) {
  if (!m.passable(to)) {
    return
  }

  from := m.pos

  m.grid.Set(from, visited)
  m.grid.Set(to, active)

  m.path = append(m.path, from)
  m.pos = to

  m.erectWalls()
}

func (m * model) passable(p gruid.Point) bool {
  return m.grid.Contains(p)
}

func (m *model) erectWalls() {
  north := m.pos.Shift(0, -1)
  east := m.pos.Shift(1, 0)
  south := m.pos.Shift(0, 1)
  west := m.pos.Shift(-1, 0)

  if m.grid.At(north).Rune == empty.Rune {
    m.grid.Set(north, wall)
  }

  if m.grid.At(east).Rune == empty.Rune {
    m.grid.Set(east, wall)
  }

  if m.grid.At(south).Rune == empty.Rune {
    m.grid.Set(south, wall)
  }

  if m.grid.At(west).Rune == empty.Rune {
    m.grid.Set(west, wall)
  }
}

func (m * model) undo(p gruid.Point) {
  last := m.path[len(m.path)-1]

  m.grid.Set(m.pos, empty)
  m.grid.Set(last, active)

  m.pos = last

  m.erectWalls()
}
