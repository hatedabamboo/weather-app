FROM golang:1.19-alpine AS builder

RUN apk update && apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY app.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM scratch

COPY --from=builder /app/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]
