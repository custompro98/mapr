package main

import (
	"context"
	"custompro98/mapr/pkg/grid"
	"fmt"
	"log"
)

func main() {
	app := grid.NewGrid(80, 24)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successful quit.\n")
	}
}
