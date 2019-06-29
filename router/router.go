package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-boilerplate/config"
	"golang-boilerplate/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cfg = config.GetConfig()

func InitRouter() {

	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})

	apiGroup := e.Group("/api")

	v1Group := apiGroup.Group("/v1")
	userHandler := new(handler.UserHandler)
	v1Group.GET("/users", userHandler.GetList)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%v", cfg.Port)); err != nil {
			log.Println("⇛ shutting down the server")
			log.Printf("%v\n", err.Error())
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
