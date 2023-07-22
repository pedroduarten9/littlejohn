# Overview

This file will have the guidelines to use the application

## Pre requisites

Have Docker installed and running.

# How to start the server

In order to start the server a helper file was created `run.sh`. This file builds the image and runs the container with the application.  
It accepts a flag `-b` that when defined builds the image before executing the container (first invocation needs this flag).

## How to interact with the application

There has been added a [postman collection](./littlejohn.postman_collection.json) to ease the exercising of the application.
If one wants to rely on curl instead, the request defintions are as follows:

### Tickers
```
curl --location 'http://localhost:8000/tickers' \
--header 'Authorization: Basic <auth>'
```

### Stock History
```
curl --location 'http://localhost:8000/tickers/<stock>/history' \
--header 'Authorization: Basic <auth>'
```

## How to check the logs

In order to check the logs of the application one can run `inspect.sh`. Once executed it will be continuously listen to the logs on the container, to stop it you should send a `SIGNINT` signal, to do this just do `CTRL + C`.

## How to stop the server

In order to stop the server a helper file was created `teardown.sh`. This executable stops the container and deletes the image.  
It accepts the flag `-i` to delete the image.

## How to add new behaviour

To add new behaviour to the API edit the [littlejohn](../littlejohn.yaml) file, run `go generate ./...` and implement the new created methods.
