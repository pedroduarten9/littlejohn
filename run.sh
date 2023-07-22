#!/bin/bash

while getopts ":b" opt; do
    case ${opt} in
        b)
        build=true
        ;;
        \?)
        echo "Invalid option: -$OPTARG" 1>&2
        exit 1
        ;;
    esac
done

if [ "$build" = true ]; then
    docker build -t littlejohn .
fi

docker run -p 127.0.0.1:8000:8000 --name littlejohn -d littlejohn