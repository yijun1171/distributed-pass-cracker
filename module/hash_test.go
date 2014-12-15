package module

import "testing"

func TestGetHash(t *testing.T) {
	hash := NewHash()
	hash.Init([]byte("hahahaha"), nil)
	t.Log(hash.GetHash())
}

func TestCheck(t *testing.T) {
	Hash := NewHash()
	src := []byte("AAAAA")
	Hash.Init(src, nil)
	t.Log("init with:", src)
	t.Log("check:", Hash.Check(src, nil))

	diff := []byte("BBBBB")
	t.Log("check with diff:", Hash.Check(diff, nil))
}
