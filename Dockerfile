FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/server ./app/cmd/server

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/migrate ./app/cmd/migrate

FROM alpine:latest

COPY --from=builder /bin/server /bin/server
COPY --from=builder /bin/migrate /bin/migrate


COPY --from=builder /build/app/migrations /app/migrations


COPY --from=builder /go/bin/goose /bin/goose

EXPOSE 8080

CMD ["/bin/server"]