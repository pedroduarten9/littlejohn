# Overview

This document will document the product next steps.

## API version

This product should have a versioned API in order to have more flexibility when providing the API to clients.

## Historical prices

The historical prices endpoint should be paginated.

## CI/CD

The application runs tests on the dockerfile in order to prevent errored builds. In the future this step should be migrated to a CI/CD pipeline.
Also the application should have a pipeline running tests and deploying frequently to ensure it's health, it also should be assessed if nightly runs for tests would prove valuable, I would say no if this is deployed frequently, but yes otherwise.

## Stock prices

It would be nice if stocks didn't vary as much as they are now, in order to tackle this I should adapt Ticker struct and enrich it with lower and upper bound.

##  Integration with real service

The first step I would take would be integrate this application with a service to give me the results instead of generating them.

## Monitoring

There should exist some monitoring such as alerts and dashboards to check for the application's health. Examples of these alerts could be like 1 5xx error, a dashboard with the activity of the system would also be valuable to assess the architecture, should this be a microservice or a lambda.

## Testing

Testing should be improved with load testing to ensure the application could stand high traffic.

## Architecture

A diagram with the architecture of the application would be very valuable to onboard new people, as it stands and no interactions exist it is no needed.