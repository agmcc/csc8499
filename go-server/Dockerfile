FROM golang:1.16
WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
 
FROM alpine:3.13.5  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/main .
EXPOSE 8080
CMD ["./main"]
