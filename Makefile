BINARY_NAME=turn

build:
 GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin turn.go
 GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux turn.go
 GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows turn.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows
