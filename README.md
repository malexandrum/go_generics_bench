# Benchmark generics func vs non-generics in Go 1.18.1 #
```
go test -bench=.
```

## On my machine ##

```
go test -bench=.
goos: darwin
goarch: amd64
pkg: example/generics
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkNonGenerics-12           463201              2660 ns/op
BenchmarkGenericsFloat-12         462338              2572 ns/op
PASS
ok      example/generics        3.584s
```

## Conclusion ##
GO Generics :D
