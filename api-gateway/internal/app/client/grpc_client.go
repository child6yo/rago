package client

// GRPClient содержит в себе клиенты для всех нижележащих сервисов, доступных по gRPC.
type GRPClient struct {
	User
}

// NewGRPCClient создает новый экземпляр GRPClient.
func NewGRPCClient(userHost string, userPort string) *GRPClient {
	return &GRPClient{
		User: *newUserClient(userHost, userPort),
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
