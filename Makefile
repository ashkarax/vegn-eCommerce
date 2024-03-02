GO=go

run:
	${GO} run ./cmd

build:
	${GO} build -o ./cmd/vegnExecutableFile ./cmd/main.go

buildrun:
	./cmd/vegnExecutableFile

swaggo:
	swag init -g ./cmd/main.go
	
swaggoformat:
	swag fmt	

test:
	${GO} test -v ./...
		