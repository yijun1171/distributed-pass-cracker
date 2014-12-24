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

func (r *RequestClient) getId() string {
	return r.clientInfo.Id
}

func (r *RequestClient) setId(id string) {
	r.clientInfo.Id = id
}

func NewReqClient(host string, port int) *RequestClient {
	clientInfo := NewClient(module.Address{host, port})
	hash := module.NewHash()
	hash.InitRand()
	return &RequestClient{*clientInfo, "", "", *hash}
}

//向服务器发出RPC调用
func (r *RequestClient) request(methodName string, args *module.Protocol, reply *module.Protocol) bool {
	err := r.clientInfo.rpcClient.Call(methodName, args, reply)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

//向服务器发出命令:HASH
func (r *RequestClient) ReqHash() bool {
	args := module.NewMsg("", r.getId(), r.hash.GetHash())
	reply := module.NewMsg("", "", "")
	result := r.request(module.MHASH, args, reply)
	log.Println("get reply from server:", reply.Commond)
	if result {
		//服务器接受任务请求
		if reply.Commond == module.MACK_JOB {
			r.setId(reply.ClientId) //save the specificed id
			return true
		}
	}
	return false
}

func (r *RequestClient) ReqPing() (bool, string) {
	args := module.NewMsg("", r.getId(), module.MPING)
	reply := module.NewMsg("", r.getId(), "")
	result := r.request(module.MPING, args, reply)
	if result {
		commond := reply.Commond
		log.Println("get ping reply from server:", commond)
		switch commond {
		case module.MNOT_DONE: //not done
			return false, ""
		case module.MDONE_NOT_FOUND: //not found
			return false, "-1"
		case module.MDONE_FOUND: //found
			return true, reply.Msg
		}
	}
	return false, ""
}
