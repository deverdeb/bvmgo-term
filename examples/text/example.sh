#!/bin/sh

go build

result=$(./text.exe)

echo "result=${result}"