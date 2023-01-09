.PHONY: env clean lint build

all: env clean lint build

env:
	@echo "=========install goimports and golint ==========="
	GOPROXY=https://goproxy.cn/,direct go install golang.org/x/tools/cmd/goimports@latest
	#install github.com/incu6us/goimports-reviser v3.3.3 by bin 
	GOPROXY=https://goproxy.cn/,direct go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

docker:
	docker rm -f octaveserver
	docker build -t octaveserver:v1.1  .
	docker run --restart=always -itd \
	-p 8630:8630 --name octaveserver octaveserver:v1.1

push:
	docker tag octaveserver:v1.1  1.117.192.82:8666/gnuoctave/octaveserver:v1.1
	docker push 1.117.192.82:8666/gnuoctave/octaveserver:v1.1
	docker rm -f octaveserver
	docker run --restart=always -itd \
	-p 8630:8630 --name octaveserver harbor.yuansuan.cn/gnuoctave/octaveserver:v1.1

build:
	go mod tidy
	gofmt -w .
	goimports -w .
	goimports-reviser -rm-unused -set-alias -format ./...
	CGO_ENABLED=0 go build  -v .

lint:
	golangci-lint run ./...

test:
	go test ./...

clean:
	go clean -i .

run:
	./octaveserver