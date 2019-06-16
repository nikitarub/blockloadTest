package main

import (
	"fmt"

	tarantool "github.com/tarantool/go-tarantool"
)

// FillIterTarantool -
func FillIterTarantool(conn *tarantool.Connection, n uint) {

	var index uint
	for index = 0; index < n; index++ {
		resp, err := conn.Insert(
			"blocks",
			[]interface{}{
				index,
				0,
				"t",
				0,
				index,
			},
		)

		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Code: ", resp.Code)
		}
		// fmt.Println("Result: ", resp)
	}
}

// UpdateIterTarantool -
func UpdateIterTarantool(conn *tarantool.Connection, n uint) {

	var index uint
	for index = 0; index < n; index++ {
		resp, err := conn.Update(
			"blocks",
			"ID",
			[]interface{}{index},
			[]interface{}{
				[]interface{}{
					"=", 2, "t",
				},
			},
		)

		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Code: ", resp.Code)
		}
		// fmt.Println("Result: ", resp)
	}
}

// DeleteIterTarantool -
func DeleteIterTarantool(conn *tarantool.Connection, n uint) {
	var index uint
	for index = 0; index < n; index++ {
		resp, err := conn.Delete(
			"blocks",
			"ID",
			[]interface{}{
				index,
			},
		)

		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Code: ", resp.Code)
		}
		// fmt.Println("Result: ", resp)
	}
}

// GetIterTarantool -
func GetIterTarantool(conn *tarantool.Connection, n uint) {
	var index uint
	for index = 0; index < n; index++ {
		resp, err := conn.Select("blocks", "ID", 0, 1, tarantool.IterEq, []interface{}{index})
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Code: ", resp.Code)
		}
		// fmt.Println("Result: ", resp)
	}
}
