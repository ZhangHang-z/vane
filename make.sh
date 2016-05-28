#!/usr/bin/env bash

vanecli='./bin/vane-cli.go'

echo "installing vane package manager"

if ! hash go; then
    echo "error: make sure you have golang installed"
else
    if [ ! `go install "$vanecli"` ]; then
                mv "$GOBIN/vane-cli" "/usr/bin/vane"
                exit 0 
        else
                echo "error: install vane failed"
exit 1    