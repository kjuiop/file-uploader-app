package main

import (
	"file-uploader-app/config"
	"file-uploader-app/internal/server"
	"log"
)

func main() {

	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("fail to read config : %s\n", err.Error())
	}

	srv := server.SetupGinServer(cfg.Server)
	srv.Run()
}
