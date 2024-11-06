#!/bin/bash

echo "Compiling..."
go build -v -ldflags "-s -w"
echo "Done compiling"

./mobile-controller-udp-server