#!/bin/bash

echo "Running test, this take a while..."

input=$(find "tests" \( -name "*_test.go" \))

for line in $input
do
    echo "------------------------------------------------------------"
    echo "Running test: $line"

    go test -v -bench=. -coverpkg=./... -coverprofile=cover.out -covermode=atomic $line
    
    echo "------------------------------------------------------------"
    echo "Finished test: $line"
done
