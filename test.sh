#!/bin/bash

go build -o ./start ./cmd/start
go build -o ./stop ./cmd/stop
go build -o ./installer ./cmd/installer

./start installer

echo
echo "Test without installer"
echo


go build -o ./start ./cmd/start
go build -o ./stop ./cmd/stop

./start
