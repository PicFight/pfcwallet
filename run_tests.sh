#!/usr/bin/env bash

GO111MODULE=on

  go version
  go clean -testcache
  go build -v ./...
  go test -v ./...
