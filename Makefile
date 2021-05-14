bench:
	@go clean -testcache ; go test -bench=. . -cpu 1,4,8