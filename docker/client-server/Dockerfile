# Build environment
# -----------------
FROM golang:1.13-alpine as build-env
WORKDIR /client-server

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/client-server .

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash

COPY --from=build-env /client-server/bin/client-server /client-server/

EXPOSE 8080
CMD ["/client-server/client-server"]