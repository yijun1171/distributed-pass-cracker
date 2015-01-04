package server

import (
	"strconv"
	"sync/atomic"
)

var count int32 = 0

func IdGenerate() string {
	id := atomic.AddInt32(&count, 1)
	return strconv.Itoa(int(id))
}
