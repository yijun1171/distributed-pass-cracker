package client

//client common struct and method
import (
	"github.com/yijun1171/Lab1/net"
)

type Client struct {
	serverAdd net.Address //server address
	id        string      //generate by server
	category  bool        //true:request_client false:worker_client
}
