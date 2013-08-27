GOPATH := $(shell pwd)

kmimg: main.go device.go
	GOPATH=$(GOPATH) go build -o $@ $^
