# Advent of Code 2023

Learn more: https://adventofcode.com/2023

## TypeScript

```shell
bun run 01/1a.ts
```

## Python

```shell
bun run 02/2a.py
```

## Golang

```shell
go run ./04/go --input ./04/input-04.txt --part A

# Build and run (much faster)
go build -o ./04/go/main ./04/go/main.go
./04/go/main --input ./04/input-04.txt --part A
```

### Bootstrap

To bootstrap Golang solution for day 5 for example

1. Copy directory `./templates/go` into `./05`
2. Update `./05/go/go.mod` file module name to `module 05`
3. Run command `go work use 05/go`
