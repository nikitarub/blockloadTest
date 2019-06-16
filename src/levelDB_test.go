package main

import (
	"fmt"
	"testing"

	ldb "technopark/blockloadTest/src/levelDB"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	NTimes = 1000
	NBatch = 100
)

func BenchmarkDeleteBatchLevelDB(b *testing.B) {
	// b.SetBytes(2)

	db, err := leveldb.OpenFile("db/test.db", nil)

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var index uint
	for index = 0; index < NTimes; index++ {
		ldb.FillLevelDBBatch(db, NBatch, index)
	}

	b.ResetTimer()

	ldb.DeleteLevelDBIter(db, NBatch*NTimes, 0)
}

func BenchmarkFillBatchLevelDB(b *testing.B) {
	// b.SetBytes(2)

	db, err := leveldb.OpenFile("db/test.db", nil)

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var index uint
	ldb.DeleteLevelDBIter(db, NBatch*NTimes, 0)

	b.ResetTimer()

	for index = 0; index < NTimes; index++ {
		ldb.FillLevelDBBatch(db, NBatch, index)
	}
}

func BenchmarkFillIterLevelDB(b *testing.B) {

	// b.SetBytes(2)

	db, err := leveldb.OpenFile("db/test.db", nil)

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// var index uint
	ldb.DeleteLevelDBIter(db, NBatch*NTimes, 0)

	b.ResetTimer()
	ldb.FillLevelDBIter(db, NBatch*NTimes, 0)

}
