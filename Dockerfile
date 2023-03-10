FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /go/bin/hazard-halt

FROM scratch

WORKDIR /

COPY domains.json ./
COPY --from=builder /go/bin/hazard-halt .

ENTRYPOINT ["/hazard-halt"]
