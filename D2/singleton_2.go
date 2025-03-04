package main

import (
	"fmt"
	"sync"
)

type DB2 struct {
	Conn string
}

var once sync.Once
var singleDB2 *DB2

func getInstance2() *DB2 {
	once.Do(func() {
		singleDB2 = &DB2{Conn: "postgresql:5432"} // instansiasi
		fmt.Println("Database instance terbuat")
	})

	return singleDB2
}

func main2() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("getInstance-", i)
			_ = getInstance2()
		}
	}()

	go func() {
		fmt.Println("goroutine ke-2")
		_ = getInstance2()
	}()

	go func() {
		fmt.Println("goroutine ke-3")
		_ = getInstance2()
	}()

	go func() {
		fmt.Println("goroutine ke-4")
		_ = getInstance2()
	}()

	go func() {
		fmt.Println("goroutine ke-5")
		_ = getInstance2()
	}()

	fmt.Scanln()
}
