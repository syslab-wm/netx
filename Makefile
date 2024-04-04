all:  vet


vet: fmt
	go vet ./...

fmt:
	go fmt ./...

# -count=1 forces tests to always run, even if no code has changed
test:
	go test -v -vet=all -count=1 ./...

.PHONY: all vet fmt
