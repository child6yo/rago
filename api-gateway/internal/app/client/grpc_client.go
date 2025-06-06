package client

import "github.com/child6yo/rago/api-gateway/internal/config"

// GRPClient содержит в себе клиенты для всех нижележащих сервисов, доступных по gRPC.
type GRPClient struct {
	User
}

// NewGRPCClient создает новый экземпляр GRPClient.
func NewGRPCClient(cfg config.Config) *GRPClient {
	return &GRPClient{
		User: *newUserClient(cfg.UserGRPCHost, cfg.UserGRPCPort),
	}
}

// StartGRPCClient устанавливает соединение со всеми содержащимися сервисами.
func (c *GRPClient) StartGRPCClient() {
	c.User.startUserClient()
}

// StartGRPCClient разрывает соединение со всеми содержащимися сервисами.
func (c *GRPClient) StopGRPCClient() {
	c.User.stopUserClient()
}
