package main

import (
	"github.com/anaseto/gruid"
	tc "github.com/gdamore/tcell/v2"
)

const (
	ColorHeader gruid.Color = 1 + iota // skip zero value ColorDefault
	ColorActive
	ColorAltBg
	ColorTitle
	ColorKey
)

type styler struct{}

func (sty styler) GetStyle(st gruid.Style) tc.Style {
	ts := tc.StyleDefault
	switch st.Fg {
	case ColorHeader:
		ts = ts.Foreground(tc.ColorNavy)
	case ColorKey:
		ts = ts.Foreground(tc.ColorGreen)
	case ColorActive, ColorTitle:
		ts = ts.Foreground(tc.ColorOlive)
	}
	switch st.Bg {
	case ColorAltBg:
		ts = ts.Background(tc.ColorBlack)
	}
	return ts
}
