package main

import (
	"fmt"
	"github.com/yijun1171/Lab1/module"
)

func main() {
	hash := module.NewHash()
	hash.Init([]byte("AAAAAA"), nil)
	fmt.Print(hash.GetHash())
}
