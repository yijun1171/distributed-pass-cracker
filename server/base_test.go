package server

import (
	"testing"
)

func TestG(t *testing.T) {
	id := IdGenerate()
	t.Logf("id string :%s", id)
}

//多线程测试原子自增
func TestIdGenerate(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			id := IdGenerate()
			t.Logf("get id %s", id)
		}()
	}
}
