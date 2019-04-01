# Episode 2 Golang GraphQL Backend
This is a simple backend for an e-commerce platform with API endpoint implemented in GraphQL
This server is built on top of the [graphql-go implmentation](https://github.com/graphql-go/graphql)

## Getting Started

### Prerequisite
* [Go](https://golang.org/doc/install) -- Development environment
* [dep](https://golang.github.io/dep/docs/installation.html) -- Dependency management
* [Docker](https://www.docker.com/products/docker-desktop)
* [Docker Compose](https://docs.docker.com/compose/install/)

### Installation
#### Setting up the environment
```sh
go get -u github.com/shredx/ep2-golang-graphql-backend
cd $GOPATH/github.com/shredx/ep2-golang-graphql-backend
dep ensure
docker-compose up
```
#### Configuring a database
 To know the name of the container run the following command
 ```sh
 docker ps
 ```
 As per docker compose configuration name of the docker container should be `ep2-golang-graphql-backend_ecommerce-mysql_1`

Login into the mysql created in the docker.
```sh
docker exec -it ep2-golang-graphql-backend_ecommerce-mysql_1 mysql -uroot -proot
```
Now create the database called `ecommerce`
```sql
create database ecommerce;
```

### Usage
Follow the [usage doc](./Usage.md)