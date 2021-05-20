# fetcher_web
A web side to serve news crawl micro services

# gRPC
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/fetchnews/v1/fetchnews.proto
```
