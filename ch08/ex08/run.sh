#!/bin/bash

go build -o go_reverb reverb1/reverb.go
./go_reverb &
go run netcat.go

