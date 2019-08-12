# Boilerplate GRPC

---

Master -> [![build status](https://github.com/technical-assessment/iqbal/salestock/grpc-inventory/badges/master/build.svg)](https://github.com/technical-assessment/iqbal/salestock/grpc-inventory/commits/master) [![coverage report](https://github.com/technical-assessment/iqbal/salestock/grpc-inventory/badges/master/coverage.svg)](https://github.com/technical-assessment/iqbal/salestock/grpc-inventory/commits/master) | Develop -> [![build status](https://github.com/technical-assessment/iqbal/salestock/grpc-inventory/badges/develop/build.svg)](https://github.com/technical-assessment/iqbal/salestock/grpc-inventory/commits/develop) [![coverage report](https://github.com/technical-assessment/iqbal/salestock/grpc-inventory/badges/develop/coverage.svg)](https://github.com/technical-assessment/iqbal/salestock/grpc-inventory/commits/develop)

---

## Requirement

- [Go 1.12 or higher](https://golang.org/dl/)

## How To Trial

- Clone this repo

```bash
git clone https://github.com/technical-assessment/iqbal/salestock/grpc-inventory.git
```

- Put outside gopath

## Regenerate Proto

```shell
protoc --go_out=plugins=grpc:. ./module/pb/inventory.proto
```
