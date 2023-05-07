# Build: docker build --no-cache --tag hello .
# Run:   docker run --name hello -p 127.0.0.1:8080:8080/tcp hello

FROM golang:alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o service . && echo "Build Success"

FROM scratch

COPY --from=builder /app /app

WORKDIR /app

EXPOSE 80

ENTRYPOINT [ "./service" ]
