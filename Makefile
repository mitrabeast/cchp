build:
	@GOGC=off go build  -o ./bin/cchp .
	@./bin/cchp

bench:
	@go test -bench=. -benchmem -v cchp/logic
