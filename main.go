package main

import (
	"context"
	"fmt"

	"github.com/erabxes/orders-api-microservices/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
		return
	}

}
