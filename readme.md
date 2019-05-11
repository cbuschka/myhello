# myhello - a hello docker container

## usage

```
docker run --rm --name myhello -p 8080:8080 cbuschka/myhello
```

or locally

```
docker build -t myhello:latest . && docker run --rm --name myhello -p 8080:8080 myhello:latest
```

## license

MIT
