package main

import (
	"advancedproject/config"
	_ "advancedproject/docs"
	"advancedproject/router"
	"advancedproject/util"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	util.BuildLogger(config.GetBaseConfig())
}

// @title API doc
// @version 1.0
// @description advanced project

// @contact.name wanghaoxi3000
// @contact.email wanghaoxi3000@163.com

// @BasePath /api/v1

// https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
func main() {
	r := router.SetupRouter()

	srv := http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			util.Log().Errorf("server start error: %s", err)
		}
	}()

	util.Log().Info("Setup router complete, run in port 3000")

	// 服务优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	util.Log().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		util.Log().Fatal("Server Shutdown:", err)
	}
	util.Log().Info("Server exist")
}
