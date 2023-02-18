FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -trimpath -o app

# ----------------------------------------------------------------

FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=builder /app/app .

ENV PORT=8080

CMD ["./app"]

# ----------------------------------------------------------------

FROM golang:1.19 as dev
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

ENV PORT=8080

CMD ["air"]
