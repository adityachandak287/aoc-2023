#!/bin/bash

echo "Building golang binary for day 05"
go build -o 05/go/main 05/go/main.go

hyperfine \
 --min-runs 1 \
 --parameter-scan maxprocs 1 8 \
 "05/go/main --input 05/input-05.txt --part BV3 --maxprocs {maxprocs}"
