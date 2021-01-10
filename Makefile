.PHONY: clean check docker test

docker:
	docker-compose run --service-ports app bash

# building for WASM
WASMDIR=wasm
GOWASM=test.wasm letterstest.wasm

DEPS=pkg/letters/*.go

ALL=$(GOWASM)
all: $(ALL)

%.wasm: export GOOS=js
%.wasm: export GOARCH=wasm
%.wasm: $(WASMDIR)/*/%.go $(DEPS)
	go build -o $@ $<

server: cmd/server.go wasm_exec.js all
	go run cmd/server.go

wasm_exec.js: /usr/local/go/misc/wasm/wasm_exec.js
	cp /usr/local/go/misc/wasm/wasm_exec.js ./

clean:
	rm -f *.wasm
	rm -f server

test: export GOOS=js
test: export GOARCH=wasm
test:
	go test -v ./...

check: export GOOS=js
check: export GOARCH=wasm
check: test
	golint ./...
	staticcheck ./...
