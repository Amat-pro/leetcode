package cache

import (
	"fmt"
	"testing"
)

func Test_LRUCache(t *testing.T) {
	c1 := New(2)
	fmt.Println("init =============================")
	c1.Print()
	c1.Put(1, 1)
	fmt.Println("after put 1 =============================")
	c1.Print()
	c1.Put(2, 2)
	fmt.Println("after put 2 =============================")
	c1.Print()
	fmt.Println("get 1: ", c1.Get(1))
	fmt.Println("after get 1 =============================")
	c1.Print()
	c1.Put(3, 3)
	fmt.Println("after put 3 =============================")
	c1.Print()
	fmt.Println("get 2: ", c1.Get(2))
	fmt.Println("after get 2 =============================")
	c1.Print()
	c1.Put(4, 4)
	fmt.Println("after put 4 =============================")
	c1.Print()
	fmt.Println("get 1: ", c1.Get(1))
	fmt.Println("after get 1 =============================")
	c1.Print()
	fmt.Println("get 3: ", c1.Get(3))
	fmt.Println("after get 3 =============================")
	c1.Print()
	fmt.Println("get 4: ", c1.Get(4))
	fmt.Println("after get 4 =============================")
	c1.Print()
}