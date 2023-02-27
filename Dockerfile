## Build
FROM golang as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY internal/ ./internal/
COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

RUN go build -o /server cmd/server/main.go

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /server /server

EXPOSE 8080

ENTRYPOINT ["/server"]