#!/bin/bash

export GOPATH=$(pwd)"/vendored";
if [ "$1" == "keys" ]; then
    ./gen_key.sh;
    go generate;
fi;
go build;
