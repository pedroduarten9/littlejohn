#
# Build stage
#
FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /api cmd/api.go

#
# Test stage
#
FROM build-stage AS run-test-stage
RUN go test -v ./...

#
# Release stage
#
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /api /api

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["/api"]