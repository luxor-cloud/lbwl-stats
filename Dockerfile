FROM golang:1.14 as builder

WORKDIR /build
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GO111MODULE=on make build

FROM alpine:3.11
WORKDIR /
RUN apk add --no-cache ca-certificates
COPY --from=builder /build/statsserver .

ENTRYPOINT ["/statsserver"]

