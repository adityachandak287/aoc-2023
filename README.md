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
> bash bench.sh

Benchmark 1: bun-ts
  Time (mean ± σ):      29.6 ms ±   3.7 ms    [User: 25.3 ms, System: 11.4 ms]
  Range (min … max):    27.2 ms …  60.4 ms    96 runs

Benchmark 2: python
  Time (mean ± σ):     118.1 ms ±   2.1 ms    [User: 69.7 ms, System: 32.3 ms]
  Range (min … max):   113.6 ms … 122.2 ms    24 runs

Benchmark 3: go-run
  Time (mean ± σ):     121.0 ms ±   4.8 ms    [User: 143.8 ms, System: 55.4 ms]
  Range (min … max):   115.4 ms … 139.3 ms    24 runs

Benchmark 4: go-binary
  Time (mean ± σ):       2.8 ms ±   0.9 ms    [User: 2.4 ms, System: 0.9 ms]
  Range (min … max):     1.9 ms …  10.0 ms    798 runs

Summary
  'go-binary' ran
   10.52 ± 3.56 times faster than 'bun-ts'
   41.97 ± 13.21 times faster than 'python'
   43.04 ± 13.64 times faster than 'go-run'
```
