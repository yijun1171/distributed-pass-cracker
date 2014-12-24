package server

import (
	"errors"
	"github.com/yijun1171/Lab1/module"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ServerHandler struct {
	clientSet module.HashSet //用于保存所有注册的客户端

}

//Hash 命令处理方法
//1.请求客户端注册 2.运算请求入队
func (h *ServerHandler) MHash(req *module.Protocol, resp *module.Protocol) error {
	log.Println("recived : commond:", req.Commond)
	if req.Commond == "" {
		return errors.New("commond is empty")
	}
	resp.Commond = module.MHASH
	return nil
}

//Ping handler
func (h *ServerHandler) MPing(req *module.Protocol, resp *module.Protocol) error {

	return nil
}

//request join handler
func (h *ServerHandler) MREQ_JOIN(req *module.Protocol, resp *module.Protocol) error {

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
