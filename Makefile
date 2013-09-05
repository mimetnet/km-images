GOPATH := $(shell pwd)

kmimg: main.go device.go
	GOPATH=$(GOPATH) go build -o $@ $^

test:
	GOPATH=$(GOPATH) go test -v

deps:
	GOPATH=$(GOPATH) go get github.com/PuerkitoBio/goquery
