package server

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"net/http"
	"syscall"
)

func Serve() {
	if err := httpRun(); err != nil {
		fmt.Printf("Http server start err: %v\n", err)
	}
	fmt.Printf("Quit server!\n")
}

func httpRun() error {
	app := gin.New()
	appRouter := SetupRouter(app)
	srv := endless.NewServer(":8976", appRouter)
	srv.BeforeBegin = func(add string) {
		fmt.Printf("Actual pid is %v\n", syscall.Getpid())
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}