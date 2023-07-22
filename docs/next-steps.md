# Overview

This document will document the product next steps.

## API version

This product should have a versioned API in order to have more flexibility when providing the API to clients.

## Historical prices

The historical prices endpoint should be paginated.

## CI/CD

The application runs tests on the dockerfile in order to prevent errored builds. In the future this step should be migrated to a CI/CD pipeline.

## Stock prices

It would be nice if stocks didn't vary as much as they are now, in order to tackle this I should adapt Ticker struct and enrich it with lower and upper bound.