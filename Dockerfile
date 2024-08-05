# syntax=docker/dockerfile:1

FROM golang:1.21.11 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags timetzdata -o /usr/local/bin/outward ./cmd/outward

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/local/bin/outward /outward

EXPOSE 8080
ENTRYPOINT ["/outward"]
