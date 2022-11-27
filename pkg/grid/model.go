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

type model struct {
	display gruid.Grid
	mapr    *gruid.Grid
	pos     gruid.Point
	bearing direction
	history []history
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
