FROM golang:alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/api-gateway-client
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make tidy build

FROM scratch
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/api-gateway-client/api-gateway-client /api-gateway-client
ENTRYPOINT ["/api-gateway-client"]
CMD []
