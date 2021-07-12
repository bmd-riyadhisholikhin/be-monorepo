# Shark

## Build and run service
If include GRPC handler, run this command (must install `protoc` compiler min version `libprotoc 3.14.0`):

```
$ make proto
```

If using SQL database, run this command for migration:
```
$ make migration dbconn="(YOUR DATABASE URL CONNECTION)"
```

And then, build and run this service:
```
$ make run
```

## Run unit test & calculate code coverage
```
$ make test
```

## Create docker image
```
$ make docker
```
