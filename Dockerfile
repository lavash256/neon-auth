FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY src/ src/

RUN go build -o main src/cmd/main.go

WORKDIR /dist
RUN cp /build/main .

FROM scratch
COPY --from=builder /dist/main /
COPY config/ config

ENTRYPOINT [ "/main", "-config","config/config.yaml"]