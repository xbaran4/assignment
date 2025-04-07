# Assignment

## Run without building
```shell
go run cmd/assignment/main.go
```

## Building
```shell
go build -o bin/assignment cmd/assignment/main.go
```

## Run built binary
```shell
./bin/assignment
```

## Build Docker image
```shell
docker build -t assignment .
```

## Run as Docker container
```shell
docker run -p 8080:8080 assignment
```

## Testing
### Create user
```shell
curl -X POST http://localhost:8080/save \
     -H "Content-Type: application/json" \
     -d '{ "external_id": "31178118-b1f7-4b6c-ab21-818b92b56714", "name": "some name", "email": "email@email.com", "date_of_birth": "2020-01-01T12:12:34+00:00" }'
```
should return:
```json
{"id":1,"external_id":"31178118-b1f7-4b6c-ab21-818b92b56714","name":"some name","email":"email@email.com","date_of_birth":"2020-01-01T12:12:34Z"}
```

### Get user
```shell
curl http://localhost:8080/1 -v 
```
should return:
```json
{"external_id":"31178118-b1f7-4b6c-ab21-818b92b56714","name":"some name","email":"email@email.com","date_of_birth":"2020-01-01T12:12:34Z"}
```
Getting user which does not exist should return 404.