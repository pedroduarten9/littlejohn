# Overview

This file will have the decisions made around the behaviour of the product

## API version

In order to follow the product strictly I chose not to version the API.

## Money format

Money was defined on the API based on [this discussion](https://github.com/OAI/OpenAPI-Specification/issues/889) and [this answer](https://github.com/OAI/OpenAPI-Specification/issues/889#issuecomment-1239978342).

## API handler generation

In order to keep things tidy I decided to create the openAPI file and generate the golang boilerplate from there.

## Go server framework

The server framework of choice was Echo, I wanted to go for the net/http package using vanilla but with so many validations I wanted to do it would become a longer, less readable application, therefore I've chosen Echo and also assessed Gin, which, between the 2, Echo had a more readable documentation, hence the decision.

## Dockerize app

The obvious choice was to dockerize the app and to make the image as light as possible the Dockerfile is a multistage one.

## Git commits

The git history will consist of small commits each with it's own responsability, not only to improve readability for for reversability concerns. The exception was one of the first commits where I wanted to guarantee the application was running and not commit an errored openAPI file.

## Middlewares

### Logging

The logging decision boiled down to the most popular logger, zap. Logging will be limited to the HTTP layer and error scenarios.

### Request ID

A request ID was added to easily integrate and debug with other systems, hence improving the traceability and identification of the request.

### Basic auth

Authentation does only check for if the username is set and password is mandatory to be empty, if one of this requirements is violated the user will be unauthorized.

## Generation

### Tickers choice

To choose tickers I choose to have only username as the seed, this has tradeoffs as the user will keep the same seeds forever, it could be interesting to add some more logic here like add the month to the seed so that the user changes tickers every month, other option would be to apply probabilistic and they had like x% chance of changing tickers every day or month. This behavious would be very easy to achieve.

### Stock price history choice

The stock price history will be generated based on the stock and the date, this way the same stock with the same date will always have the same price.

## Pagination

As required I added pagination with just the page, it could be achieved with an offset and limit, or page and limit as well.
I limited to 40 weeks max which is approximately 10 years.
Sending no page is the same as sending page = 1.

## Validation

The validations were done in code as there was just one on the pagination, if there was more it would be good to tag the structs and use the validator from golang to validate structs.

## Others

ExistsTicker is concrete to Tickers for the sake of simplicity, could be generalized to every type of slice easily.