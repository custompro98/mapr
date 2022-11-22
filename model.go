package main

import (
	"github.com/anaseto/gruid"
)

var EMPTY = gruid.Cell{Rune: ' '}
var VISITED = gruid.Cell{Rune: '.'}
var ACTIVE = gruid.Cell{Rune: 'O'}
var WALL = gruid.Cell{Rune: 'X'}

type model struct {
  grid gruid.Grid
  pos gruid.Point
}

func NewModel(gd gruid.Grid) *model {
	m := &model{
		grid: gd,
    pos: gruid.Point{
      X: 0,
      Y: gd.Ug.Height / 2,
    },
	}

	return m
}

func (m *model) InitializeMap() {
  m.Move(m.pos)
}

func (m *model) Update(msg gruid.Msg) gruid.Effect {
	switch msg := msg.(type) {
	case gruid.MsgInit:
		m.InitializeMap()
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
	case gruid.KeyEscape, "Q", "q":
		return gruid.End()
  default:
    return nil
	}

  m.Move(delta)

	return nil
}

func (m *model) Draw() gruid.Grid {
	return m.grid
}

func (m *model) Move(to gruid.Point) {
  if (!m.passable(to)) {
    return
  }

  from := m.pos

  m.grid.Set(from, VISITED)
  m.grid.Set(to, ACTIVE)

  m.pos = to

  m.erectWalls()
}

func (m *model) erectWalls() {
  north := m.pos.Shift(0, -1)
  east := m.pos.Shift(1, 0)
  south := m.pos.Shift(0, 1)
  west := m.pos.Shift(-1, 0)

  if m.grid.At(north).Rune == EMPTY.Rune {
    m.grid.Set(north, WALL)
  }

  if m.grid.At(east).Rune == EMPTY.Rune {
    m.grid.Set(east, WALL)
  }

  if m.grid.At(south).Rune == EMPTY.Rune {
    m.grid.Set(south, WALL)
  }

  if m.grid.At(west).Rune == EMPTY.Rune {
    m.grid.Set(west, WALL)
  }
}

func (m * model) passable(p gruid.Point) bool {
  return m.grid.Contains(p)
}
