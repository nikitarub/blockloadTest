# blockloadTest
Testing db's for load with data blocks

# Results
In `test_results.txt`

# benchmark levelDB
## from src/ 
```bash
sh doTest.sh
```

# benchmark Tarantool
## setup 
Required scripts in `src/tarantool/app.lua`

## from src/tarantool
```bash
sh doTest.sh
```

# Tune bench params
## doTest.sh
```
-count 4 # amount of test runs
-cpu 4 
-benchtime 3s # max time for each run
```

## levelDB_test.go
```go
const (
	NTimes = 1000 // how many batches to add
    NBatch = 100 // batches 
    // or just NTimes * NBatch = NSize --> amount of blocks to add
)
```

## tarantool_test.go
```go
const (
	nInsert = 1000 // amount of blocks to add
)
```
