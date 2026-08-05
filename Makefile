.PHONY: test expect-fail build

test:
	go test ./...

build:
	go build -o bin/radixd ./cmd/radixd

expect-fail:
	@if go test ./internal/resp -count=1 >/dev/null 2>&1; then echo "unexpected pass"; exit 1; fi
	@echo "expect-fail ok: resp tests still red"
