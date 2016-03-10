#!/bin/sh

cd /vagrant
go get ./...
nohup go run service.go >/vagrant/app.log 2>&1 &
