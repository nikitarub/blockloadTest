package main

import (
	"fmt"
	b "technopark/blockloadTest/src/block"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/vmihailenco/msgpack"
)

// var loginFormTmpl = []byte(`
// <html>
// 	<body>
// 	<form action="/login" method="post">
// 		Login: <input type="text" name="login">
// 		Password: <input type="password" name="password">
// 		<input type="submit" value="Login">
// 	</form>
// 	</body>
// </html>
// `)

// func innerPage(w http.ResponseWriter, r *http.Request) {

// 	db, err := leveldb.OpenFile("db/test.db", nil)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()

// 	iter := db.NewIterator(nil, nil)
// 	for iter.Next() {
// 		// key := iter.Key()
// 		value := iter.Value()

// 		var valueUnmarshal b.Block
// 		err := msgpack.Unmarshal(value, &valueUnmarshal)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// fmt.Printf("key: %d | value: %s\n", valueUnmarshal.ID, valueUnmarshal.Data)
// 		str := []byte(valueUnmarshal.Data + "\n")
// 		w.Write(str)
// 	}

// 	// w.Write(loginFormTmpl)
// 	w.Header().Set("Content-Type", "plain-text")
// }

func main() {
	// ––––––––––––––––––––––––––––––––––––––––
	// ––––––––––––––––––––––––––––––––––––––––
	// ––––––––––––––––––––––––––––––––––––––––

	db, err := leveldb.OpenFile("db/test.db", nil)

	fmt.Println(err)
	defer db.Close()

	// fillLevelDBBatch(db, 10)
	// fillLevelDBIter(db, 10)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// key := iter.Key()
		value := iter.Value()

		var valueUnmarshal b.Block
		err = msgpack.Unmarshal(value, &valueUnmarshal)
		if err != nil {
			panic(err)
		}

		fmt.Printf("key: %d | value: %s\n", valueUnmarshal.ID, valueUnmarshal.Data)
	}

	// var stat *leveldb.DBStats
	// err = db.Stats(stat)
	// if err != nil {
	// 	panic(err)
	// }

	// http.HandleFunc("/", innerPage)

	// fmt.Println("starting server at :8080")
	// http.ListenAndServe(":8080", nil)
}
