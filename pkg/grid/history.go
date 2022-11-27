package grid

import "github.com/anaseto/gruid"

type point struct {
  point gruid.Point
  cell gruid.Cell
}

type history struct {
  points []point
  pos gruid.Point
  bearing direction
}

func (m *model) snapshot() {
  h := history {
    pos: m.pos,
    bearing: m.bearing,
    points:make([]point, 0),
  }

  it := m.mapr.Iterator()
  for it.Next() {
    h.points = append(h.points, point{
      point: it.P(),
      cell: it.Cell(),
    })
  }

  m.history = append(m.history, h)
}

func (m *model) undo() {
	if len(m.history) == 1 {
		return
	}

	last := m.history[len(m.history)-1]
  m.history = m.history[:len(m.history)-1]
  m.pos = last.pos
  m.bearing = last.bearing

  for _, v := range last.points {
    m.mapr.Set(v.point, v.cell)
  }
}
