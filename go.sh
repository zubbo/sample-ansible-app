#!/bin/sh

lsof -t -i tcp:80 -s tcp:listen | xargs kill -9
cd /vagrant
go get ./...
nohup go run service.go >/vagrant/app.log 2>&1 &
