package module

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

var table []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

type Hash struct {
	hashStr string
	hashCtx hash.Hash
}

func NewHash() *Hash {
	return &Hash{"", md5.New()}
}

func (h *Hash) Init(src []byte, salt []byte) string {
	h.hashCtx.Write(src)
	h.hashStr = hex.EncodeToString(h.hashCtx.Sum(salt))
	return h.GetHash()
}

func (h *Hash) GetHash() string {
	return h.hashStr
}

func (h *Hash) Check(other []byte, salt []byte) bool {
	if other == nil {
		return false
	}
	h.hashCtx.Reset() //reset
	h.hashCtx.Write(other)
	otherStr := hex.EncodeToString(h.hashCtx.Sum(salt))
	if otherStr == h.hashStr {
		return true
	} else {
		return false
	}
}

//范围检测, 检测成功返回true和对应值,失败返回false和空串
func (h *Hash) CheckRange(start string, end string) string {
	if !less(start, end) { //start > end illegalParameters
		return ""
	}
	var current = start
	var last = end
	for current != last {
		if h.Check([]byte(current), nil) { //不加盐
			return current
		} else {
			current = Next(current)
		}
	}
	if h.Check([]byte(last), nil) {
		return last
	} else {
		return ""
	}
}

//判断两个第一个参数是否小于等于第二个参数
func less(left string, right string) bool {
	sliceL := []byte(left)
	sliceR := []byte(right)
	for key, value := range sliceL {
		if getIndex(value) > getIndex(sliceR[key]) {
			return false
		}
	}
	return true
}

func getIndex(v byte) int {
	for key, value := range table {
		if value == v {
			return key
		}
	}
	return -1
}

//range AAAAA ~ 99999
//超过最大范围时返回空串
func Next(current string) string {
	if current == "99999" {
		return ""
	}
	cur := []byte(current)
	curLen := len(cur)
	for i := curLen - 1; i >= 0; i-- {
		for key, value := range table {
			if value == cur[i] {
				nextIndex := (key + 1) % len(table)
				cur[i] = table[nextIndex] //更新值
				if nextIndex != 0 {       //是否进位
					return string(cur)
				} else {
					break
				}
			}
		}
	}
	return ""
}
