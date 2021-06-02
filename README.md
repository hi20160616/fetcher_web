# Fetch News
A web side to serve news crawl micro services

# gRPC

## grpc-gateway
Refer:  
https://github.com/grpc-ecosystem/grpc-gateway  
https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/  
https://www.cnblogs.com/FireworksEasyCool/p/12782137.html  

## fix import "google/api/annotations.proto"  error
### Copy google/api/annotations.proto and google/api/http.proto
Refer:  
https://github.com/grpc-ecosystem/grpc-gateway/issues/1935  
```
mkdir -p google/api
wget https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -O google/api/annotations.proto
wget https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -O google/api/http.proto
```
## Generate
```
protoc -I "./" -I "../../../" \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/fetchnews/web/v1/fetchnews.proto
go mod tidy
```

# TODO
- [ ] add: content display
- [ ] add: search implement
- [ ] fix: dumplicate double loop range for objects type switch
- [ ] change: handler path: list to ms title, such as bbc, voa, etc.
- [ ] optimize: update skeleton to deal better for biz, data, service
