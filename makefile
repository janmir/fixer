all: build
	#use make build to build for developement
	#make build-release for release version
	./dev

init:
	sls create -t aws-go-dep -p service

build:
	go build -race -tags dev -o dev ./fixer

build-release:
	dep ensure
	env GOOS=linux go build -tags release -ldflags="-s -w" -o bin/fixer ./src

remove:
	sls remove

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build-release
	sls deploy --verbose
