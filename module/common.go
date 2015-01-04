package module

import (
	"fmt"
	"log"
	//"strings"
)

//basic net structure and method

//rpc调用方法名常量
const (
	MHASH           = "MHash"
	MACK_JOB        = "AckJob"
	MREQ_JOIN       = "MReqJoin"
	MJOB            = "MJob"
	MPING           = "MPing"
	MDONE_NOT_FOUND = "MDoneNotFound"
	MDONE_FOUND     = "MDoneFound"
	MNOT_DONE       = "MNotDone"
)

type Address struct {
	IP   string
	Port int
}

//format IP:port
func (a *Address) GetHost() string {
	//return strings.Join([]string{a.IP, ":", string(a.Port)}, "")
	return fmt.Sprintf(a.IP+":%d", a.Port)
}

//msg format
type Protocol struct {
	Commond  string
	ClientId string
	Msg      string //msg content
}

func NewMsg(commond string, clientId string, msg string) *Protocol {
	return &Protocol{commond, clientId, msg}
}

//组合service和method
func GetMethod(name string) string {
	return "ServerHandler." + name
}

func CheckFatal(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
