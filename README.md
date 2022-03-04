[toc]


# Postgres

## 安装
> sudo mkdir -p /data/postgres/data  
> sudo docker run --name postgres -v /data/postgres/data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=test@2022 -p 15432:5432 -d postgres

## 命令行操作
> sudo docker exec -it postgres bash
```
> psql -U postgres -p 5432
postgres=# create user htsample with password 'htsample';
postgres=# create database htsample owner htsample;
postgres=# grant all on htsample to htsample;
```

## 连接相关

> psql -h dev.vincent78.top -p 15432 -U htsample -d htsample

> jdbc:postgresql://dev.vincent78.top:15432/htsample?sslmode=disable

## 初始化语句
[init.sql](doc/db/init.sql)