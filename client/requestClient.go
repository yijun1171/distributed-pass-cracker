package client

import (
	"github.com/yijun1171/Lab1/module"
	"log"
)

//request client
type RequestClient struct {
	clientInfo Client
	ciphertext string //密文
	plaintext  string //明文
	hash       module.Hash
}

func (r RequestClient) getId() string {
	return r.clientInfo.Id
}

func NewReqClient(host string, port int) *RequestClient {
	clientInfo := NewClient(module.Address{host, port})
	hash := module.NewHash()
	hash.InitRand()
	return &RequestClient{*clientInfo, "", "", *hash}
}

//向服务器发出RPC调用
func (r RequestClient) request(methodName string, args *module.Protocol, reply *module.Protocol) bool {
	err := r.clientInfo.rpcClient.Call(methodName, args, reply)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

//向服务器发出命令:HASH
func (r RequestClient) ReqHash() bool {
	args := module.NewMsg("", r.getId(), r.hash.GetHash())
	reply := module.NewMsg("fd", "", "")
	result := r.request(module.MHASH, args, reply)
	log.Println("get reply from server:", reply.Commond)
	if result {
		if reply.Commond == module.MACK_JOB {
			return true
		}
	}
	return false
}
