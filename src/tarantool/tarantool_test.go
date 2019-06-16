package main

import (
	"fmt"
	"testing"

	// tt "technopark/blockloadTest/src/tarantool"

	tarantool "github.com/tarantool/go-tarantool"
)

const (
	nInsert = 10000
)

func BenchmarkDeleteIterTarantool(b *testing.B) {

	// b.SetBytes(2)

	opts := tarantool.Opts{User: "guest"}
	conn, err := tarantool.Connect("127.0.0.1:3301", opts)
	// conn, err := tarantool.Connect("/path/to/tarantool.socket", opts)
	if err != nil {
		fmt.Println("Connection refused:", err)
	}
	FillIterTarantool(conn, nInsert)
	b.ResetTimer()

	DeleteIterTarantool(conn, nInsert)
}

func BenchmarkFillIterTarantool(b *testing.B) {
	opts := tarantool.Opts{User: "guest"}
	conn, err := tarantool.Connect("127.0.0.1:3301", opts)
	// conn, err := tarantool.Connect("/path/to/tarantool.socket", opts)
	if err != nil {
		fmt.Println("Connection refused:", err)
	}
	DeleteIterTarantool(conn, nInsert)
	b.ResetTimer()

	FillIterTarantool(conn, nInsert)
}
