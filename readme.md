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

## How to Run Locally
```shell
make run
```
or simply with go command
```shell
go run main.go
```

## Endpoint URL
1. `GET localhost:1234/api/v1/articles?page=0&size=1&query=place` get list article
curl
```shell
curl --location 'localhost:1234/api/v1/articles?page=0&size=1&query=place'
```

2. `GET localhost:1234/api/v1/articles/12` Get Detail Article
curl
```shell
curl --location 'localhost:1234/api/v1/articles/12'
```

3. `POST localhost:1234/api/v1/articles` create article
curl
```shell
curl --location 'localhost:1234/api/v1/articles' \
--header 'Content-Type: application/json' \
--data '{
    "author": "radiohead",
    "title": "jigsaw falling into place",
    "body": "<div>Just as you take my hand<br>Just as you write my number down<br>Just as the drinks arrive<br>Just as they play your favourite song<br>As your bad day disappears<br>No longer wound up like a spring<br>Before you'\''ve had too much<br>Come back in focus again<br>"
}'
```