package server

import (
	"context"
	"file-uploader-app/config"
	"sync"
)

type MockServer struct {
	cfg config.Server
}

func NewMockClient(cfg config.Server) Client {
	return &MockServer{
		cfg: cfg,
	}
}

func (m MockServer) Run(wg *sync.WaitGroup) {
	panic("implement me")
}

func (m MockServer) Shutdown(ctx context.Context) {
	panic("implement me")
}
