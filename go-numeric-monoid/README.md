# go-numeric-monoid

This package is a Go binding of [NumericMonoid](https://github.com/hivert/NumericMonoid) used by [DANSE](https://github.com/efournival/ter-lri/tree/master/danse).

Architecture support is currently limited to `amd64`.

## Building

Building this package requires a GCC version having Cilk++ support (typically all GCC releases starting from 4.9).

It also needs an additionnal preprocessor flag defining the `MAX_GENUS` to be computed.

For example:
```
CGO_CPPFLAGS='-DMAX_GENUS=35' go test -v
CGO_CPPFLAGS='-DMAX_GENUS=35' go get -u 'github.com/efournival/ter-lri/go-numeric-monoid'
```
