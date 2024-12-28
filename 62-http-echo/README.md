## Postgres 

1. postgres

```

```


2. adminer (DB UI)


## Using docker

- Go to https://hub.docker.com/ and signup for a docker account(if you dont have an account )

```
docker login 
```
- create a docker network

```
docker network create ps-network
```

- container for postgres

```
docker run -d --name demo-pg --network ps-network -p
 45432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin123 -e POSTGRES_DB=usersdb post
gres
```

- database UI

```
docker run -d --name ps-db-ui -p 38080:8080 --network ps-network adminer 
```

- To build a docker image

```
docker build . -t jpalaparthi/primesoft-demo:v0.1
```

```
docker run -d --name ps-app1 --network ps-network -e
 PORT=9093 -e DB_CON="host=demo-pg user=admin password=admin123 dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai" -p 9093:9093 jpalaparthi/primesoft-demo:v0.1
 ```
 
