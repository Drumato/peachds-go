# peachds - Simple Data Structures in Go -

![Go Version](https://img.shields.io/github/go-mod/go-version/drumato/peachds-go)
![Build Status](https://github.com/drumato/peachds-go/actions/workflows/ci.yml/badge.svg)
![Coverage Status](https://coveralls.io/repos/github/drumato/peachds-go/badge.svg?branch=main)
![License](https://img.shields.io/github/license/drumato/peachds-go)

## Features

- OrderedMap

## Installation

To install the project, use `go get`:

```sh
go get github.com/drumato/peachds-go

```


## Benchmark Result

```text
goos: linux
goarch: amd64
pkg: github.com/Drumato/peachds-go
cpu: 13th Gen Intel(R) Core(TM) i7-13700F
BenchmarkConcurrentOrderedMap_Set
BenchmarkConcurrentOrderedMap_Set-24             9637089               152.7 ns/op           114 B/op          0 allocs/op
BenchmarkConcurrentOrderedMap_Get
BenchmarkConcurrentOrderedMap_Get-24            30595050                59.56 ns/op            0 B/op          0 allocs/op
BenchmarkConcurrentOrderedMap_Iter
BenchmarkConcurrentOrderedMap_Iter-24           15483910               104.2 ns/op             0 B/op          0 allocs/op
BenchmarkOrderedMap_Set
BenchmarkOrderedMap_Set-24                      14501823               141.8 ns/op           139 B/op          0 allocs/op
BenchmarkOrderedMap_Get
BenchmarkOrderedMap_Get-24                      46411224                45.13 ns/op            0 B/op          0 allocs/op
BenchmarkOrderedMap_Iter
BenchmarkOrderedMap_Iter-24                     15537765               105.4 ns/op             0 B/op          0 allocs/op
```

