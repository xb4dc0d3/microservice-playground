# Stage 1
FROM golang:1.21.1-alpine AS builder

ARG GIT_SHA

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go /app

RUN CGO_ENABLED=0 go build -ldflags "-X main.sha=${GIT_SHA}" -o /app/go-app

# Stage 2
FROM alpine:3.19

COPY --from=builder /app .

EXPOSE 5000

CMD ["./go-app"]