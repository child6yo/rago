package client

// type Storage struct {
// 	splitter pb.SplitterServiceClient
// 	conn     *grpc.ClientConn

// 	host string
// 	port string
// }

// func newStorageClient(host string, port string) *Storage {
// 	return &Storage{host: host, port: port}
// }

// func (sc *Storage) startStorageClient() {
// 	addr := fmt.Sprintf("%s:%s", sc.host, sc.port)
// 	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Print("failed to connect splitter grpc server")
// 	}

// 	sc.splitter = pb.NewSplitterServiceClient(conn)
// 	sc.conn = conn
// }

// func (sc *Storage) stopStorageClient() {
// 	if sc.conn == nil {
// 		return
// 	}
// 	sc.conn.Close()
// }
