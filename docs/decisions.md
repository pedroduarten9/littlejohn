# Overview

This file will have the decisions made around the behaviour of the product

## API version

In order to follow the product strictly I chose not to version the API.

## Money format

Money was defined on the API based on [this discussion](https://github.com/OAI/OpenAPI-Specification/issues/889) and [this answer](https://github.com/OAI/OpenAPI-Specification/issues/889#issuecomment-1239978342).

## API handler generation

In order to keep things tidy I decided to create the openAPI file and generate the golang boilerplate from there.
