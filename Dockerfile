FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /src
WORKDIR /src

RUN go build -o bin/neon-auth cmd/neon-auth/main.go 
RUN go build -o bin/neon-migrate cmd/migrate/main.go

WORKDIR /dist
RUN cp /src/bin/neon-auth .
RUN cp /src/bin/neon-migrate .

FROM alpine:3.4
COPY --from=builder /dist/neon-auth /bin
COPY --from=builder /dist/neon-migrate /bin
COPY configs/ config/
COPY migrations/ migrations/
COPY tools/entrypoint.sh .

RUN chmod 777 entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

