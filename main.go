package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/test_cache_CQRS/boot"
	"github.com/test_cache_CQRS/config"
	"github.com/test_cache_CQRS/router"
)

func main() {
	flag.StringVar(&config.Env, "env", "local", "A config name that used by server")
	flag.Parse()

	setup := boot.MakeHandler()
	handlerRouter := router.NewHandlerRouter(setup)
	app := handlerRouter.RouterWithMiddleware()

	port := fmt.Sprintf(":%v", config.Conf.Port)
	if port == "" {
		port = fmt.Sprintf(":%v", 1234)
	}

	log.Printf("Server running on port %s", port)
	serve := &http.Server{
		Addr:    port,
		Handler: app,
	}

	// Start server
	go func() {
		if err := app.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 1 second.
	quit := make(chan os.Signal)
	// kill (no param) default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := serve.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
	log.Println("Server exiting")
}
