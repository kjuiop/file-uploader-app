package main

import (
	"context"
	"file-uploader-app/config"
	"file-uploader-app/internal/server"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("fail to read config : %s\n", err.Error())
	}

	srv := server.NewGinServer(cfg.Server)

	wg.Add(1)
	go srv.Run(&wg)

	<-exitSignal()
	srv.Shutdown(ctx)
	cancel()
	wg.Wait()
	log.Println("server gracefully stopped")
}

func exitSignal() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}
