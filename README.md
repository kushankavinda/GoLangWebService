## How to RUN
1. Install go 1.20 version
2. go run .\cmd\main.go
3. How to build in docker
    docker build -t {application-name}:{version} .
    docker push {application-name}:{version}



## Remove unused Go Modules
go mod tidy

## Go Lang Access Modifiers
First Letter of the attribute should be capital to access in the outside

## Go generics