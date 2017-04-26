#!/bin/sh

rm -rf $GOPATH/pkg/linux_amd64/github.com/efournival/ter-lri/go-numeric-monoid.a $GOPATH/src/github.com/efournival/ter-lri
CGO_CPPFLAGS='-DMAX_GENUS=42' go get "github.com/efournival/ter-lri/go-numeric-monoid"
