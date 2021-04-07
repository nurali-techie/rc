build: clean
	go build -o ./out/ ./cmd/rc

clean:
	rm -f ./out/rc

go-mod:
	go mod tidy
	go mod verify

test-integration:
	go test ./integration -v
