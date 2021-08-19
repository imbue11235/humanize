build:
	go build ./...

cover:
	go test -coverprofile=cover.out -covermode=atomic -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

test:
	go test ./... -v