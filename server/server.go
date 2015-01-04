package server

import (
	"container/list"
	"errors"
	"github.com/yijun1171/Lab1/module"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ServerHandler struct {
	clientSet  module.HashSet //用于保存所有注册的客户端
	workerSet  module.HashSet //保存所有的worker
	taskQueue  list.List
	clientChan chan clientInfo //用于传输客户端的注册和注销
	taskChan   chan string     //用于传输运算请求
}

const (
	ACTION_REGISTER bool = true
	ACTION_ERASE    bool = false
)

type clientInfo struct {
	id     string
	action bool
}

//远程调用方法注册对象
func NewServerHandler() *ServerHandler {
	handler := ServerHandler{}
	handler.clientSet = *module.NewHashSet()
	handler.taskQueue = *list.New()
	handler.clientChan = make(chan clientInfo, 10)
	handler.taskChan = make(chan string, 10)
	handler.init()
	return &handler
}

func (h *ServerHandler) init() {
	//监听注册和注销信息
	go func() {
		for {
			select {
			case info := <-h.clientChan:
				if info.action == ACTION_REGISTER { //注册
					log.Println("register client-id :", info.id)
					h.clientSet.Add(info.id)
				} else if info.action == ACTION_ERASE { //注销
					log.Println("remove client-id:", info.id)
					h.clientSet.Remove(info.id)
				}
			case task := <-h.taskChan: //接收运算请求
				if task != "" {
					log.Println("task enqueue :", task)
					h.taskQueue.PushBack(task) //入队
				}
			}
		}
	}()
}

//return generate id
func (h *ServerHandler) addClient() string {
	id := IdGenerate()
	h.clientChan <- clientInfo{id, ACTION_REGISTER}
	return id
}

//Hash 命令处理方法
//1.请求客户端注册 2.运算请求入队
func (h *ServerHandler) MHash(req *module.Protocol, resp *module.Protocol) error {
	log.Println("recived : commond:", req.Commond)
	if req.Commond == "" {
		return errors.New("commond is empty")
	}
	//客户端注册
	resp.Commond = module.MACK_JOB //回复ack_job
	resp.ClientId = h.addClient()  //返回生成的id
	h.taskChan <- req.Msg          //将请求入队
	return nil
}

//Ping handler
func (h *ServerHandler) MPing(req *module.Protocol, resp *module.Protocol) error {

	return nil
}

//request join handler
func (h *ServerHandler) MREQ_JOIN(req *module.Protocol, resp *module.Protocol) error {
	log.Println("recived : commond:", req.Commond)
	if req.Commond == "" {
		return errors.New("commond is empty")
	}
	//客户端注册
	resp.Commond = module.MACK_JOB //回复ack_job
	resp.ClientId = h.addClient()  //返回生成的id
	h.taskChan <- req.Msg          //将请求入队
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

	serverHandler := NewServerHandler()
	server.Register(serverHandler)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("server recevied")
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
