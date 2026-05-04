.PHONY: format utest test

format:
	gofumpt -l -w $$(go list -f '{{.Dir}}' ./...)

utest:
	go test ./...

test: utest
