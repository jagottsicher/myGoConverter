BINARY_NAME=turn

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/macos/${BINARY_NAME} ./src/turn.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/linux/${BINARY_NAME} ./src/turn.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/windows/${BINARY_NAME}.exe ./src/turn.go

run:
	./bin/linux/${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ./bin/macos/${BINARY_NAME}
	rm ./bin/linux/${BINARY_NAME}
	rm ./bin/windows/${BINARY_NAME}.exe
