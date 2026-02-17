FROM golang:1.25.7-bookworm AS builder

WORKDIR /app

# copy packages files
COPY go.* ./
RUN go mod download

# copy source files
COPY ./ ./

RUN go build -v -o poc_motor ./src/


FROM debian:bookworm-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/poc_motor /app/poc_motor

CMD ["/app/poc_motor"]
