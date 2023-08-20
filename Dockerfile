FROM golang:alpine AS builder
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/hello-webapp /app/
COPY files/ /app/files
WORKDIR /app
CMD ["./hello-webapp"]
