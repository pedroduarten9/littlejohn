#!/bin/bash

while getopts ":i" opt; do
    case ${opt} in
        i)
        deleteImage=true
        ;;
        \?)
        echo "Invalid option: -$OPTARG" 1>&2
        exit 1
        ;;
    esac
done


docker stop littlejohn
docker rm littlejohn

if [ "$deleteImage" = true ]; then
    docker image rm littlejohn
fi