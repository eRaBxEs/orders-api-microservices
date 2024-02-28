package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{}),
	}
	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %v", err)
	}

	defer func() { // to be executed whenever the Start function ends
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server") // added to see that everything is going on well

	ch := make(chan error, 1) // buffer size of int 1: in an unbuffered channel the writer/publisher will always be blocked
	// whenever they write to a channel until that value is read off of a channel

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}

		close(ch) // closes the channel when the function is done,
		// closing a channel also sends a signal to anyone listening
		// that the channel is done informing the readerr to stop expecting data on it.

	}()

	// Now to set up a receiver for our channel
	//err = <-ch // captures any value sent into our channel into our error variable
	// we can also use a second return value which determines the current state of the channel, open if true or closed if false
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10) // giving enought time of 10 seconds for any inflight request to be done
		defer cancel()
		return server.Shutdown(timeout)
	}

}
