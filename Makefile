GOPATH := $(shell pwd)

kmimg: main.go device.go
	GOPATH=$(GOPATH) go build -o $@ $^

deps:
	GOPATH=$(GOPATH) go get github.com/PuerkitoBio/goquery
