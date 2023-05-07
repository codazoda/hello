# Information for Publishing

## Uploading to dockerhub

You can run this image from dockerhub. I use the following commands to push it there.

```
docker tag hello:latest codazoda/hello:latest
docker push codazoda/hello:latest
```

Explanation:

- The `docker tag` command renames the local `hello:latest` image to `codazoda/hello:latest`.
- The `docker push` command pushes the image up to dockerhub.
