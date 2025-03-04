package main

import (
	"fmt"
)

type DB struct {
	Conn string
}

var singleDB *DB

func getInstance() *DB {
	if singleDB == nil {
		singleDB = &DB{Conn: "postgresql:5432"} // instansiasi
		fmt.Println("Database instance terbuat")
	}

	return singleDB
}

func main1() {
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
