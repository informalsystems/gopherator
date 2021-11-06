# gopherator [![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/informalsystems/gopherator.svg)](https://github.com/informalsystems/gopherator) [![Go Report Card](https://goreportcard.com/badge/github.com/informalsystems/gopherator)](https://goreportcard.com/report/github.com/informalsystems/gopherator) [![Go Reference](https://pkg.go.dev/badge/github.com/informalsystems/gopherator.svg)](https://pkg.go.dev/github.com/informalsystems/gopherator)
[Modelator](https://github.com/informalsystems/modelator)'s cousin for Golang

---
### Example
[Golang port](https://github.com/informalsystems/gopherator/tree/main/examples/numbersystem) of [NumberSystem](https://github.com/informalsystems/modelator/blob/main/modelator/tests/integration/resource/numbers.rs)

#### Instruction
```sh
git clone git@github.com/informalsystems/gopherator
cd gopherator
cd third_party/mbt
cargo build --release
cd -
cd examples/numbersystem
go test -v
```

##### Output
```sh
=== RUN   TestFixedExecutions
=== RUN   TestFixedExecutions/test_0
=== RUN   TestFixedExecutions/test_1
--- PASS: TestFixedExecutions (0.00s)
    --- PASS: TestFixedExecutions/test_0 (0.00s)
    --- PASS: TestFixedExecutions/test_1 (0.00s)
=== RUN   TestModelBased
2021/11/04 02:43:48 Generating traces using Modelator (rs-binding)...
=== RUN   TestModelBased/[test:_AMaxBMaxTest,_trace:_0]
=== RUN   TestModelBased/[test:_AMaxBMinTest,_trace:_0]
=== RUN   TestModelBased/[test:_AMinBMaxTest,_trace:_0]
--- PASS: TestModelBased (2.81s)
    --- PASS: TestModelBased/[test:_AMaxBMaxTest,_trace:_0] (0.00s)
    --- PASS: TestModelBased/[test:_AMaxBMinTest,_trace:_0] (0.00s)
    --- PASS: TestModelBased/[test:_AMinBMaxTest,_trace:_0] (0.00s)
PASS
ok  	github.com/informalsystems/gopherator/examples/numbersystem	2.809s
```
