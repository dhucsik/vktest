package main

import (
	"context"
	"log"

	"github.com/dhucsik/vktest/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.InitApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
