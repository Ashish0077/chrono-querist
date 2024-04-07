FROM golang:1.22-alpine3.19 AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build cmd/web/main.go

FROM alpine
RUN apk add tzdata
ENV TZ=Asia/Kolkata

WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]