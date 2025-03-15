package temp

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var temp int
var lock sync.Mutex

func AddTemp() {
	lock.Lock()
	temp++
	lock.Unlock()
}

// 两个goroutine交替打印，然后自增一个全局变量
func PrintGouroutine(count int) {

	chanA := make(chan struct{})
	chanB := make(chan struct{})
	chanC := make(chan struct{})

	go PrintA(chanA, chanB)
	go PrintB(chanB, chanC)
	go PrintC(chanC, chanA, count)

	chanA <- struct{}{}

	time.Sleep(10 * time.Second)

}

func PrintA(chanA chan struct{}, chanB chan struct{}) {

	for _ = range chanA {
		fmt.Println("A")
		AddTemp()
		chanB <- struct{}{}
	}

}

func PrintB(chanB, chanC chan struct{}) {
	for _ = range chanB {
		fmt.Println("B")
		AddTemp()
		chanC <- struct{}{}
	}
}

func PrintC(chanC, chanA chan struct{}, n int) {
	count := 0
	for _ = range chanC {
		fmt.Println("C")
		AddTemp()
		count++
		if count == n {
			fmt.Println("tempValue is: ", temp)
			break
		}
		chanA <- struct{}{}
	}
}

func Test_PrintGouroutine(t *testing.T) {
	PrintGouroutine(10)
}
