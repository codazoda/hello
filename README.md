# Hello

A very simple go http server used for testing docker configurations.

## Run in a container

I've pushed this code to dockerhub so you can simply pull the image and run it.

```
docker pull codazoda/hello:latest
docker run --name hello -p 127.0.0.1:8080:8080/tcp codazoda/hello`
```

## Other Information

[Publishing](publish.md)  
[Development](develop.md)  
