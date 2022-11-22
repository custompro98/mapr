package main

import (
	"context"
	"fmt"
	"log"

	"github.com/anaseto/gruid"
	tcell "github.com/anaseto/gruid-tcell"
)

func main() {
	grid := gruid.NewGrid(80, 24)

	driver := tcell.NewDriver(tcell.Config{
		StyleManager: styler{},
	})

	model := NewModel(grid)

	app := gruid.NewApp(gruid.AppConfig{
		Driver: driver,
		Model:  model,
	})

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successful quit.\n")
	}
}
