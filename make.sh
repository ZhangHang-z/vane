#!/usr/bin/env bash

vanecli='./bin/vane-cli.go'

echo "installing vane package manager"

if hash go; then 
	echo "error: make sure you have golang installed"
	exit 1
else
	if [ `go install "$vanecli"` ]; then
		ln -s /usr/bin/vane "$GOPATH/bin/vane-cli.exe"
	else
		echo "error: install vane failed"
	fi
fi