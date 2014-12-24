package module

import (
	"bytes"
	"fmt"
)

//用字典类型作为底层 实现set数据结构
type HashSet struct {
	m map[interface{}]bool //键类型是空接口,任何数据类型都是它的实现,所以可以存储任何值(函数,字典,切片类型除外).值类型是bool,简化代码书写和节省空间
}

func NewHashSet() *HashSet { //返回值类型是指针
	return &HashSet{m: make(map[interface{}]bool)} //用make函数进行初始化
}

func (set *HashSet) Add(e interface{}) bool { //方法接收者是当前值的复制,声明成指针类型以减小开销,并可以更改接收者的属性值
	if !set.m[e] { //不存在元素e时 将e加入set
		set.m[e] = true
		return true
	}
	return false
}

func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e) //使用delete内建函数 专门用于删除字典类型中的元素
}

func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool) //将m重新赋值为空的字典,旧的字典会被GC
}

func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

func (set *HashSet) Len() int {
	return len(set.m) //内建函数 可以返回 string [n]T []T map[K]T chan T 这些类型的长度
}

func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//用于产生set的快照
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m { //copy
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else { //迭代过程中 m中的元素可能增多
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen { //迭代过程中 m中元素可能减少 以至于slice未被填满 需要截取
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}
