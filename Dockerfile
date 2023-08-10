FROM golang:1.17-alpine as builder

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download


# intermediate stage: Build the binary
FROM golang:1.17-alpine as runner

COPY --from=builder ./app ./app

RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /app
ENV config=docker

EXPOSE 1234

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main