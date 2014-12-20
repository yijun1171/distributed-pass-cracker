package server

import (
	"github.com/yijun1171/Lab1/client"
	"testing"
	"time"
)

var requestClient *client.RequestClient

func init() {
	go startServer(8888)
	time.Sleep(time.Duration(5) * time.Second)
	requestClient = client.NewReqClient("localhost", 8888)
}

func TestRequestHash(t *testing.T) {
	requestClient.ReqHash()
	t.Log("success")
}
