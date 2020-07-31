.PHONY: build

build:
	export GO111MODULE=on
	env GOOS=linux GORACH=amd64 go build -ldflags="-s -w" -o bin/machinestream_linux_amd64 *.go
	env GOOS=darwin GORACH=amd64 go build -ldflags="-s -w" -o bin/machinestream_mac_amd64 *.go

