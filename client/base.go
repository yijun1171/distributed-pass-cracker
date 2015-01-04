package client

//client common struct and method
import (
	"github.com/yijun1171/Lab1/module"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//client base struct
type Client struct {
	ServerAdd module.Address //server address
	Id        string         //generate by server
	Category  bool           //true:request_client false:worker_client
	rpcClient *rpc.Client
}

func NewClient(add module.Address) *Client {
	host := add.GetHost()
	client, e := jsonrpc.Dial("tcp", host)
	//acts like dial, connects to the address on the named network
	//Dial("tcp", "12.34.56.78:80") address format is host:name
	//DialTimeout(network, address, timeout)
	// timeout Duration: int64 base on nanosecond
	if e != nil {
		log.Println("connect failed:", e.Error())
	}
	return &Client{add, "", true, client}
}

func (c *Client) closeClient() {
	c.rpcClient.Close()
}
