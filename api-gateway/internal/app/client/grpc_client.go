package client

import "github.com/child6yo/rago/api-gateway/internal/config"

// GRPClient содержит в себе клиенты для всех нижележащих сервисов, доступных по gRPC.
type GRPClient struct {
	*User
	*Storage
	*Generator
}

// NewGRPCClient создает новый экземпляр GRPClient.
func NewGRPCClient(cfg config.Config) *GRPClient {
	return &GRPClient{
		User: newUserClient(cfg.UserGRPCHost, cfg.UserGRPCPort),
		Storage: newStorageClient(cfg.SplitterGRPCHost, cfg.SplitterGRPCPort),
		Generator: newGeneratorClient(cfg.GeneratorGRPCHost, cfg.GeneratorGRPCPort),
	}
}

// StartGRPCClient устанавливает соединение со всеми содержащимися сервисами.
func (c *GRPClient) StartGRPCClient() {
	c.User.startUserClient()
	c.Storage.startStorageClient()
	c.Generator.startGeneratoClient()
}

// StartGRPCClient разрывает соединение со всеми содержащимися сервисами.
func (c *GRPClient) StopGRPCClient() {
	c.User.stopUserClient()
	c.Storage.stopStorageClient()
	c.Generator.stopGeneratoClient()
}
