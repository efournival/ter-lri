# go-numeric-monoid

This package is a Go binding of [NumericMonoid](https://github.com/hivert/NumericMonoid) used by [DANSE](https://github.com/efournival/ter-lri/tree/master/danse).

## Building

Due to some limitations in the original C++ code, building this package requires GCC 5.

It also needs an additionnal preprocessor flag defining the `MAX_GENUS` to be computed.

Typically:
```
CGO_CPPFLAGS='-DMAX_GENUS=35' CC=gcc-5 CXX=g++-5 go build
CGO_CPPFLAGS='-DMAX_GENUS=35' CC=gcc-5 CXX=g++-5 go test -v
```
