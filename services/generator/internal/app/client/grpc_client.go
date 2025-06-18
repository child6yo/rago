package client

import "github.com/child6yo/rago/services/generator/internal/config"

// GRPClient содержит в себе клиенты для всех нижележащих сервисов, доступных по gRPC.
type GRPClient struct {
	Storage
}

// NewGRPCClient создает новый экземпляр GRPClient.
func NewGRPCClient(cfg config.Config) *GRPClient {
	return &GRPClient{
		Storage: *newStorageClient(cfg.StorageGRPCHost, cfg.StorageGRPCPort),
	}
}

// StartGRPCClient устанавливает соединение со всеми содержащимися сервисами.
func (c *GRPClient) StartGRPCClient() {
	c.Storage.startStorageClient()
}

// StopGRPCClient разрывает соединение со всеми содержащимися сервисами.
func (c *GRPClient) StopGRPCClient() {
	c.Storage.stopStorageClient()
}
