build:
	go build -o ./bin/goblk

run: build 
	./bin/goblk

test:	
	go test -v ./...