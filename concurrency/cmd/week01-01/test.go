package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %v. \n", data)
	}

	fmt.Println("-----임계영역이 보호되는 코드 /w sync.Mutex -----")
	var memoryAccess sync.Mutex
	var value int

	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("this value is %v\n", value)
	} else {
		fmt.Printf("this value is %v\n", value)
	}
	memoryAccess.Unlock()
}
