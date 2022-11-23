package grid

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/ui"
)

var empty = gruid.Cell{Rune: ' '}
var visited = gruid.Cell{Rune: '.'}
var active = gruid.Cell{Rune: 'O'}
var wall = gruid.Cell{Rune: 'X'}

var help = []string{
  "h: move left",
  "j: move down",
  "k: move up",
  "l: move right",
  "u: undo",
  "?: help",
  "q: quit",
}

type model struct {
  display gruid.Grid
  mapr *gruid.Grid
  pos gruid.Point
  path []gruid.Point
  help *ui.Pager
}

func newModel(gd gruid.Grid) *model {
	m := &model{
    display: gd,
    mapr: &gd,
    pos: gruid.Point{
      X: 0,
      Y: gd.Ug.Height / 2,
    },
    help: newPager(PagerConfig{
      w: gd.Ug.Width,
      h: gd.Ug.Height,
      title: "Help",
      body: help,
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
  var delta = m.pos

  m.display = *m.mapr

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
  case "?":
    m.display = m.help.Draw()
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

  m.mapr.Set(from, visited)
  m.mapr.Set(to, active)

  m.path = append(m.path, from)
  m.pos = to

  m.erectWalls()
}

func (m * model) passable(p gruid.Point) bool {
  return m.mapr.Contains(p)
}

func (m *model) erectWalls() {
  north := m.pos.Shift(0, -1)
  east := m.pos.Shift(1, 0)
  south := m.pos.Shift(0, 1)
  west := m.pos.Shift(-1, 0)

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

func (m * model) undo(p gruid.Point) {
  last := m.path[len(m.path)-1]

  m.mapr.Set(m.pos, empty)
  m.mapr.Set(last, active)

  m.pos = last

  m.erectWalls()
}
