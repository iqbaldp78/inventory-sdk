FROM golang:1.12-alpine as builder
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/technical-assessment/iqbal/salestock/grpc-inventory
COPY vendor ../vendor
COPY grpc-inventory ./

RUN go build -o main main.go

FROM alpine:3.8
COPY --from=builder /go/src/github.com/technical-assessment/iqbal/salestock/grpc-inventory/main /app/
COPY --from=builder /go/src/github.com/technical-assessment/iqbal/salestock/grpc-inventory/config.json /app/
COPY --from=builder /go/src/github.com/technical-assessment/iqbal/salestock/grpc-inventory/salestock.db /app/
WORKDIR /app
EXPOSE 8010
CMD ["./main"]
