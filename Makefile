.PHONY: build clean

build:
	dep ensure -v
	env GOOS=linux go build main.go -ldflags="-s -w" -o bin/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy_prod: clean build
	sls deploy --verbose --stage production

test:
	go test -v ./...