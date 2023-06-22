.PHONY: gen build run clean

build:
	go build main.go

run:
	go run main.go

gen:
	protoc --go_out=paths=source_relative:. proto/shadow.proto

clean:
	rm proto/*.go
