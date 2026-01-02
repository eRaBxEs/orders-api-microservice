package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/erabxes/orders-api/application"
)

func main() {
	app := application.New(application.LoadConfig())

	// use context.background() to derive a root level context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
