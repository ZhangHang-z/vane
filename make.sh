#!/bin/sh

echo "installing vane package manager"

if type go 2> /dev/null; then 
	echo "error: make sure you have golang installed"
	exit 1
else
	cd "./bin"
	if go install vane-cli.go 2> /dev/null; then 
		echo "error: install vane-cli failed"
	else
		ln -s /usr/bin/vane $GOPATH/bin/vane-cli
	fi
fi