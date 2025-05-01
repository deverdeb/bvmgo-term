#!/bin/sh

go build

result=$(./select.exe "Choose a color (q to quit):" "reb" "blue" "green" "yellow" "other color")

echo "result=${result}"