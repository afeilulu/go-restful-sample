# go-gorm-swaggo
Simple REST service in Go with Gofiber, PostgreSql, GORM, SWAGGO

***
### Setup and usage

1. create database mannually
```
CREATE DATABASE example_db WITH OWNER postgres ENCODING 'utf-8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';
```
2. update .env(product environment) or .env.developemnt.local(development environment) for your own environment
```
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=example_db
USERNAME=
PASSWORD=
PREFORK=false
```

3. cd to the project directory

```go
./make run
```

or build a docker image to run

```go
./make docker
```

Create Group
```shell
curl -X 'POST' \
  'http://localhost:9090/api/groups' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "memo": "group memo",
  "name": "group name"
}'
```

Get Groups
```shell
curl http://localhost:9090/groups
```

Update Group
```shell
curl -X 'POST' \
  'http://localhost:9090/api/groups/1' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "memo": "test55555",
  "name": "test55555"
}'
```

Delete Group
```shell
curl -X DELETE http://localhost:9090/groups/1
```

***
### Refer

1. [gofiber](https://github.com/gofiber/fiber): https://docs.gofiber.io/
2. [gorm](https://gorm.io/): https://learnku.com/docs/gorm/v2
