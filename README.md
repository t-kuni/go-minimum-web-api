# go-minimum-web-api

A web api server for communication checks. The response includes information about the request. The same information is also output to the log (standard output).

```
docker run --rm -it -p 33333:80 tkuni83/go-minimum-web-api:latest
```

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"太郎", "age":"30"}' http://localhost:33333/aaa/bbb
```

# development

```
go run main.go

# or specify port
PORT=34567 go run main.go
```

```
docker build --tag tkuni83/go-minimum-web-api:latest .
docker login
docker push tkuni83/go-minimum-web-api:latest
```