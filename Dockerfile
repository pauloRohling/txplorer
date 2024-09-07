FROM golang:1.23.1 as builder
ARG CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags="-w -s" -o=./cmd/txplorer.tmp ./cmd

FROM gruebel/upx:latest as upx
ARG UPX_VERSION=4.2.4
WORKDIR /app
COPY --from=builder /app/cmd .
RUN upx --lzma --best -o ./txplorer ./txplorer.tmp
RUN rm -rf ./txplorer.tmp

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=upx /app .
EXPOSE 8080
ENTRYPOINT ["./txplorer"]