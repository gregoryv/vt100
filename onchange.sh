#!/bin/bash -e
GOPATH=/home/gregory
go test -cover -coverprofile /tmp/c.out .
uncover /tmp/c.out
