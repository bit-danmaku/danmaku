FROM golang:alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc
COPY go.mod go.sum ./
RUN GOPROXY=https://goproxy.cn,direct go mod download
COPY . .
RUN make tidy build

FROM scratch
ARG BIN_NAME
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/app/bin/$BIN_NAME /app
ENTRYPOINT ["/app"]
CMD []
