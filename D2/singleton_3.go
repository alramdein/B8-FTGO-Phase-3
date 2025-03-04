package main

import (
	"fmt"
	"sync"
	"time"
)

type DB3 struct {
	Conn string
}

var singleDB3 *DB3
var mu sync.Mutex

func getInstance3() *DB3 {
	mu.Lock()
	defer mu.Unlock()

	if singleDB3 == nil {
		time.Sleep(time.Millisecond * 3000)       // simulasi instansasi dbnya lama
		singleDB3 = &DB3{Conn: "postgresql:5432"} // instansiasi
		fmt.Println("Database instance terbuat")
	}

	return singleDB3
}

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("getInstance-", i)
			_ = getInstance3()
		}
	}()

	go func() {
		fmt.Println("goroutine ke-2")
		_ = getInstance3()
	}()

	go func() {
		fmt.Println("goroutine ke-3")
		_ = getInstance3()
	}()

	go func() {
		fmt.Println("goroutine ke-4")
		_ = getInstance3()
	}()

	go func() {
		fmt.Println("goroutine ke-5")
		_ = getInstance3()
	}()

	fmt.Scanln()
}
