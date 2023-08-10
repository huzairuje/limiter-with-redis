# README

## Description
this repository basically implementing limiter rate (Token Bucket Algorithm) and cache (CQRS pattern) 
on golang using echo framework (https://echo.labstack.com/), postgresql and redis.


## Install PreRequisite
1. run `make local` to install or up or run the postgres container and redis container.
1. make sure the docker postgres or container postgres is running and make database on the postgres container with name `test_cache_cqrs`
and run the sql script on the directory `migrations`
2. make sure redis container is up and running
3. make sure the config on config.local.yaml is correct

## How to Run Localy
```shell
make run
```
or simply with go command
```shell
go run main.go
```