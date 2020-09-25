#!/bin/bash

go build -o go_reverb reverb2/reverb.go
./go_reverb &
go run netcat.go

