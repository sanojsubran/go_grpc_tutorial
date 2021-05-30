module github.com/sanojsubran/client

go 1.16

replace github.com/sanojsubran/pipe => ../proto/gen

require (
	github.com/sanojsubran/pipe v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.38.0
)
