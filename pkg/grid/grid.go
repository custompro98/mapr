package grid

import (
	"context"

	"github.com/anaseto/gruid"
	tcell "github.com/anaseto/gruid-tcell"
)

type Grid interface {
  Start(context.Context) error
}

func NewGrid(w, h int) Grid {
	grid := gruid.NewGrid(w, h)

	driver := tcell.NewDriver(tcell.Config{
		StyleManager: styler{},
	})

	model := newModel(grid)

	app := gruid.NewApp(gruid.AppConfig{
		Driver: driver,
		Model:  model,
	})

  return app
}
