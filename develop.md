# Information for Development

## Run the go code locally

`go run hello.go`

## Run with an a port specified

`PORT=8086 go run hello.go`

## Build the executable

`go build hello.go`

## Build a docker image

`docker build --no-cache --tag codazoda/hello .`

## Run a docker container

`docker run --name hello -p 127.0.0.1:8080:8080/tcp --network=example codazoda/hello`
