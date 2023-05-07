# Hello

A very simple go http server used for testing docker configurations.

## Run in a container

I've pushed this code to dockerhub so you can simply pull the image and run it.

```
docker pull codazoda/hello:latest
docker run --name hello -p 8080:8080 codazoda/hello
```

## Other Information

[Development](develop.md)  
[Publishing](publish.md)  
