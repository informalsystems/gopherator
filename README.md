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
