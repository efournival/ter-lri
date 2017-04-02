#!/bin/sh

rm -f $GOPATH/pkg/linux_amd64/github.com/efournival/ter-lri/go-numeric-monoid.a
CGO_CPPFLAGS='-DMAX_GENUS=35' go get -u "github.com/efournival/ter-lri/go-numeric-monoid"
