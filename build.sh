#!/bin/sh

#golangci-lint -j 1 run ./... --timeout=10m
go env -w GOOS=linux
go build -p 1 -tags=middle_srv -o bin/middle_srv_rpc ./app/rpc/main.go
go build -p 1 -tags=middle_srv -o bin/middle_srv_job ./app/job/main.go


