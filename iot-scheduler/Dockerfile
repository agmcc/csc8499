FROM golang:1.16
WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o iot-scheduler cmd/main.go

FROM alpine:3.13.5
RUN apk --no-cache add ca-certificates
WORKDIR /usr/local/bin
COPY --from=0 /go/src/iot-scheduler .
CMD ["/usr/local/bin/iot-scheduler"]
