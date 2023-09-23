.PHONY: test
test:
	go test ./internal/... ./cmd/... -coverprofile cover.out

.PHONY: bin
bin: fmt vet
	go build -o /usr/local/bin/kubectl-colorize_applied ./cmd/kubectl-colorize-applied

.PHONY: fmt
fmt:
	go fmt ./internal/... ./cmd/...

.PHONY: vet
vet:
	go vet ./internal/... ./cmd/...
