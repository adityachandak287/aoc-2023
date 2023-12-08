# Advent of Code 2023

Learn more: https://adventofcode.com/2023

## Typescript

```shell
bun run 01/1a.ts

# Compile Typescript to Javascript (for node)
npx tsc 04/typescript/4b.ts --target es2016 --module nodenext --moduleResolution nodenext
node 04/typescript/4b.js
# Skip compile step with ts-node (much slower)
npx ts-node 04/typescript/4b.ts
```

## Python

```shell
python3 02/2a.py

# profile python code (output in index.html, decrease interval if not enough samples)
pyinstrument --interval=0.000001 --renderer html --outfile index.html 04/python/4b.py
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

## Comparing performance by benchmarking using [hyperfine](https://github.com/sharkdp/hyperfine)

```shell
bash 04/bench.sh
```

Benchmark output: [04/bench.log](04/bench.log)
