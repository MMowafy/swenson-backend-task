## IV. USE CASES
All espresso machines
- EM001
- EM002 
- EM003 

All small pods
- CP001 
- CP003 
- CP011
- CP013
- CP021
- CP023
- CP031 
- CP033 
- CP041 
- CP043

All pods sold in 7 dozen packs
- EP007
- EP017

## How to use

```
## install dependencies first 

go get -v  ./...

## To build and start app

docker-compose up --build 

## after postgres container is up exec inside container to create database

docker exec -it postgres sh

## inside container run 

psql -h localhost -U postgres -w -c "create database swenson_db;"
exit

## hit API

http://localhost:8080/coffee-products?productType=machine
http://localhost:8080/coffee-products?productType=machine&&waterLineCompatible=true
http://localhost:8080/coffee-products?productType=large machine

http://localhost:8080/coffee-products?productType=pod
http://localhost:8080/coffee-products?productType=pod&&packSize=12
http://localhost:8080/coffee-products?productType=pod&&packSize=12&&flavor=mocha

## to run tests
go test -v ./...
```
