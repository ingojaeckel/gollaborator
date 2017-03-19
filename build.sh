#!/bin/sh
rm -f app app.bz2

CGO_ENABLED=0 GOOS=linux go build -a -o app
file app
du -h app

bzip2 app
du -h app.bz2
