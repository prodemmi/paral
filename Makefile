antlr:
	antlr4 -Dlanguage=Go -o antlr antlr/Paral.g4

build:
	go build -o paral main.go

test-examples:
	antlr4 -Dlanguage=Go -o antlr antlr/Paral.g4
	go build -o paral main.go
	./paral examples/simple.paral
	./paral examples/print.paral
	./paral examples/ignore_errors.paral
	./paral examples/stash.paral

test:
	@go test ./... -v