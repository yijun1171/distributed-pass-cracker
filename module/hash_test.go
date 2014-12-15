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

func TestNext(t *testing.T) {
	t.Log("AAAAA's next is :", Next("AAAAA"))
	t.Log("AAAAZ's next is :", Next("AAAAZ"))
	t.Log("AAAAz's next is :", Next("AAAAz"))
	t.Log("AAAB9's next is :", Next("AAAB9"))
	t.Log("00099's next is :", Next("00099"))
	t.Log("99998's next is :", Next("99998"))
}

func TestRangeCheck(t *testing.T) {
	Hash := NewHash()
	Hash.Init([]byte("BHGSD"), nil)
	t.Log("range: BHAAA - BHGGG", Hash.CheckRange("BHAAA", "BHGGG"))
	t.Log("range: BHAAA - BH000", Hash.CheckRange("BHAAA", "BH000"))
}

func TestLess(t *testing.T) {
	t.Log("ABSHD < ABSM1", less("ABSHD", "ABSM1"))
	t.Log("F3300 < Fll00", less("F3300", "Fll00"))
}
