.PHONY: build docker lint updateDep

build:
	go build main.go

updateDep:
	go get -u=patch
	go mod vendor

lint:
	golangci-lint run ./...
	go vet ./...

docker:
	docker build -t kubeedge-database -f Dockerfile .
