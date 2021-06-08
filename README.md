# Fetch News
A web side to serve news crawl micro services

# Usage
## Add MS
Just edit https://gist.github.com/hi20160616/277faf779bee0d0d1525696eff6e8b56

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
- [x] add: content display
- [x] add: search implement
- [x] add: markdown view for content
<!-- - [ ] fix: dumplicate double loop range for objects type switch -->
- [x] fix: Miscroservice client conns close action block on Miscroservice server closed.
<!-- - [ ] change: handler path: list to ms title, such as bbc, voa, etc. -->
- [x] optimize: update skeleton to deal better for biz, data, service
- [x] optimize: Miscroservice Client Conns open should display if success.
- [ ] optimize: error display elegant while runtime err: invalid memory address
