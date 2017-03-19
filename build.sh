#!/bin/sh
CGO_ENABLED=0 GOOS=linux go build -a -o app
du -hs app
file app
