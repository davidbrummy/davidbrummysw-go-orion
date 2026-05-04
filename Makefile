.PHONY: db-down db-up format utest test

db-up:
	docker compose up -d postgres

db-down:
	docker compose down

format:
	gofumpt -l -w $$(go list -f '{{.Dir}}' ./...)

utest:
	go test ./...

test: utest
