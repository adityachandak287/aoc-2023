#!/bin/bash

DAY=04

echo "Building golang binary for day ${DAY}"
go build -o $DAY/go/main $DAY/go/main.go

hyperfine \
 --warmup 10 \
 --min-runs 10 \
 -n bun-ts "bun run 04/typescript/4b.ts" \
 -n python "python3 04/python/4b.py" \
 -n go-run "go run ./$DAY/go --input $DAY/input-$DAY.txt --part B" \
 -n go-binary "$DAY/go/main --input $DAY/input-$DAY.txt --part B"
