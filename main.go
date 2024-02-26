package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/erabxes/orders-api-microservices/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
		return
	}

	//cancel() // if we called this func earlier it will cancel our derived context and any of the children context underneath it

}
