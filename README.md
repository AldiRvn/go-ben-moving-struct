# go-ben-moving-struct
Go Benchmark of Moving Struct

## How To Run

```bash
go test -test.fullpath=true -benchmem -run=^$ -bench "^(Benchmark_normalPerField|Benchmark_marshallUnmarshall)$" go-ben-moving-struct -v -count=1
```