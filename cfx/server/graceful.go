package server

import (
	"cfx_server/warehouses/app"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Graceful() {

	engine := GetEngine()

	// 运行 http 服务
	srv := &http.Server{
		Addr:              ":" + app.Config("app").GetString("port"),
		Handler:           engine,
		ReadTimeout:       3600 * time.Second,
		WriteTimeout:      3600 * time.Second,
		ReadHeaderTimeout: 500 * time.Millisecond,
		MaxHeaderBytes:    1 << 20, // 1 MB
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
