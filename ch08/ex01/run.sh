#!/bin/bash

go build -o go_clock clock/clock.go
TZ=US/Eastern ./go_clock -port 8081 &
TZ=Asia/Tokyo ./go_clock -port 8082 &
TZ=Europe/Longon ./go_clock -port 8083 &
go run clockwall.go eastern=localhost:8081 tokyo=localhost:8082 london=localhost:8083

