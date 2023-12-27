build:
	export GIN_MODE=release
	go build main.go wire_gen.go

run-dev:
	export GIN_MODE=debug
	go run main.go wire_gen.go

run-test:
	export GIN_MODE=test
	go test -v ./test/...
