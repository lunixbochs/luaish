.PHONY: parser.go
all : parser.go

parser.go : parser.go.y
	goyacc -o $@ -v table.out parser.go.y; [ -f y.output ] && ( rm -f y.output )
