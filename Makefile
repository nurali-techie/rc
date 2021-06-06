build: clean
	go build -o ./out/rc .

install: build
	go install

clean:
	rm -f ./out/rc

go-mod:
	go mod tidy
	go mod verify

test:
	go test `go list ./... | grep -v tests`

test-integration:
	go test ./tests -v
