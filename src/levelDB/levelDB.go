package ldb

import (
	b "technopark/blockloadTest/src/block"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/vmihailenco/msgpack"
)

// DeleteLevelDBIter - удаление
func DeleteLevelDBIter(db *leveldb.DB, n uint, pow uint) {

	var index uint
	for index = n * pow; index < n*(pow+1); index++ {
		indexByte, err := msgpack.Marshal(&index)
		if err != nil {
			panic(err)
		}

		err = db.Delete([]byte(indexByte), nil)
		if err != nil {
			// fmt.Println(err)
			panic(err)
		}
	}
}

// FillLevelDBIter - итеративное заполнение
func FillLevelDBIter(db *leveldb.DB, n uint, pow uint) {

	var index uint
	for index = n * pow; index < n*(pow+1); index++ {
		var block = b.Block{
			ID:         index,
			DocumentID: 0,
			Data:       "l",
			Line:       0,
			Column:     index,
		}

		blBite, err := msgpack.Marshal(&block)
		if err != nil {
			panic(err)
		}
		indexByte, err := msgpack.Marshal(&index)
		if err != nil {
			panic(err)
		}

		err = db.Put([]byte(indexByte), blBite, nil)
		if err != nil {
			// fmt.Println(err)
			panic(err)
		}
	}
}

// FillLevelDBBatch - заполнение пачкой
func FillLevelDBBatch(db *leveldb.DB, n uint, pow uint) {
	batch := new(leveldb.Batch)

	var index uint
	for index = n * pow; index < n*(pow+1); index++ {
		var block = b.Block{
			ID:         index,
			DocumentID: 0,
			Data:       "l",
			Line:       0,
			Column:     index,
		}

		blBite, err := msgpack.Marshal(&block)
		if err != nil {
			panic(err)
		}
		indexByte, err := msgpack.Marshal(&index)
		if err != nil {
			panic(err)
		}

		batch.Put([]byte(indexByte), blBite)

	}

	err := db.Write(batch, nil)
	if err != nil {
		panic(err)
	}
}
