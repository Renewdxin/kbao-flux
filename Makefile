.PHONY: test vet run lint-openapi

test:
	go test ./...

vet:
	go vet ./...

run:
	go run ./cmd/mock-server

lint-openapi:
	postman spec lint openapi/kbao-flux.public.openapi.yaml --output json
