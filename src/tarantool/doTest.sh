go test -benchtime 3s -bench . -benchmem -count 4 -cpu 4 
# -benchsplit=10
# -cpuprofile cpu.out -memprofile mem.out