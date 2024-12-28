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
