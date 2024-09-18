FROM golang:1.22 as builder
ARG VERSION="undefined"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -a -installsuffix cgo \
    -o myapp \
    -ldflags="-X 'github.com/KingscliffHH/app/pkg/version.Version=${VERSION}'" \
    cmd/api/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/myapp .

EXPOSE 8000

ENTRYPOINT [ "./myapp"]
