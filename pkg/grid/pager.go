package grid

import (
	"github.com/anaseto/gruid"
	"github.com/anaseto/gruid/ui"
)

const (
	ColorTitle gruid.Color = 1 + iota // skip zero value ColorDefault
	ColorLnum
)

type PagerConfig struct {
  w int
  h int
  title string
  body []string
}

func newPager(cnf PagerConfig) *ui.Pager {
  lines := make([]ui.StyledText, len(cnf.body))

  for i, v := range cnf.body {
    lines[i] = ui.Text(v)
  }

  pager := ui.NewPager(ui.PagerConfig{
    Grid: gruid.NewGrid(cnf.w, cnf.h),
    Lines: lines,
    Style: ui.PagerStyle{
      LineNum: generateStyle(ColorLnum),
    },
    Box: &ui.Box{
      Style: generateStyle(ColorTitle),
      Title: ui.Text("Help"),
    },
  })

  return pager
}

func generateStyle(fg gruid.Color) gruid.Style {
  return gruid.Style{
    Fg: fg,
  }
}
