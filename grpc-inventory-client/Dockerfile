FROM golang:1.12-alpine as builder
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/technical-assessment/iqbal/salestock/grpc-inventory-client
COPY vendor ../vendor
COPY grpc-inventory-client ./
RUN go build -o main main.go

FROM alpine:3.8
COPY --from=builder /go/src/github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/main /app/
COPY --from=builder /go/src/github.com/technical-assessment/iqbal/salestock/grpc-inventory-client/config.json /app/
WORKDIR /app
EXPOSE 8080
CMD ["./main"]
