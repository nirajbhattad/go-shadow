#!/bin/bash

default="coverage.out"
coverage_path=${1:-$default}

if [ "$1" == "-coverage" ]; then
	coverage_path="coverage/coverage.out"
fi

go tool cover -html=$coverage_path -o coverage/coverage.html
if [ "$1" == "-coverage" ]; then
	gocov convert $coverage_path | gocov-html > coverage/summary.html
fi
