package grid

import (
	"github.com/anaseto/gruid"
	tc "github.com/gdamore/tcell/v2"
)

const (
	colorHeader gruid.Color = 1 + iota // skip zero value ColorDefault
	colorActive
	colorAltBg
	colorTitle
	colorKey
)

type styler struct{}

func (sty styler) GetStyle(st gruid.Style) tc.Style {
	ts := tc.StyleDefault
	switch st.Fg {
	case colorHeader:
		ts = ts.Foreground(tc.ColorNavy)
	case colorKey:
		ts = ts.Foreground(tc.ColorGreen)
	case colorActive, colorTitle:
		ts = ts.Foreground(tc.ColorOlive)
	}
	switch st.Bg {
	case colorAltBg:
		ts = ts.Background(tc.ColorBlack)
	}
	return ts
}
