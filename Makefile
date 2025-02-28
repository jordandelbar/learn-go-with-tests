.PHONY: test
test:
	@go test $(if $(v),-v) ./...

.PHONY: benchmark
benchmark:
	@cd iteration && go test -bench=.

.PHONY: doc
doc:
	@godoc -http=:8080
