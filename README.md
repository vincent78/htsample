[toc]


# Postgres

## setup the db
> sudo mkdir -p /data/postgres/data  
> sudo docker run --name postgres -v /data/postgres/data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=test@2022 -p 15432:5432 -d postgres

## operation in terminal
> sudo docker exec -it postgres bash
```
> psql -U postgres -p 5432
postgres=# create user htsample with password 'htsample';
postgres=# create database htsample owner htsample;
postgres=# grant all on htsample to htsample;
```

## connect the db 

> psql -h dev.vincent78.top -p 15432 -U htsample -d htsample

> jdbc:postgresql://dev.vincent78.top:15432/htsample?sslmode=disable

## init sql file
[init.sql](doc/db/init.sql)


# release
## build command
> make build

## the command in terminal
> build/bin/htsample --host 127.0.0.1 --port 8080 --dbHost dev.vincent78.top --dbPort 15432 --dbUser htsample --dbPwd htsample --dbName htsample server


## gen local docker （MAC)
> make docker


## gen docker for running in linux
> make dockerRelease

