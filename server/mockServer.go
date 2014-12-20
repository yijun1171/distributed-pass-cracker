package server

import (
	"github.com/yijun1171/Lab1/module"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ServerHandler struct{}

//Hash 命令处理方法
func (h *ServerHandler) MHash(req *module.Protocol, resp *module.Protocol) error {
	log.Println("recived : commond:", req.Commond)
	resp.Commond = module.MHASH
	return nil
}

func (h *ServerHandler) Hash(req *int, resp *int) error {
	return nil
}
func startServer(port int) {
	server := rpc.NewServer()

	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatal("server\t-", "listen error:", err.Error())
	}

	log.Println("server listening on port 8888")
	defer listener.Close()

	serverHandler := ServerHandler{}

	server.Register(&serverHandler)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("server recevied")
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
