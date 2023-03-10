FROM golang:1.20-bullseye AS builder

COPY go.mod .
COPY go.sum .

WORKDIR /app

RUN go mod download

COPY . .

RUN go build -o /bin/app

FROM debian:bullseye-slim

COPY --from=builder /bin/app /bin/app
COPY --from=builder /app/domains.json /bin/app/domains.json

ENTRYPOINT ["/bin/app"]
