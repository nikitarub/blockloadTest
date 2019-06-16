package main

import (
	"fmt"

	tarantool "github.com/tarantool/go-tarantool"
)

func main() {
	opts := tarantool.Opts{User: "guest"}
	conn, err := tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		fmt.Println("Connection refused:", err)
	}
	const nInsert = 100
	fmt.Println("\nDelete")
	DeleteIterTarantool(conn, nInsert)
	fmt.Println("\nInsert")
	FillIterTarantool(conn, nInsert)
	fmt.Println("\nUpdate")
	UpdateIterTarantool(conn, nInsert)
	fmt.Println("\nSelect")
	GetIterTarantool(conn, nInsert)
	fmt.Println("\nDelete")
	DeleteIterTarantool(conn, nInsert)
}
