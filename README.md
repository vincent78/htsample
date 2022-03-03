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
```

create table sys_dict_type (
    id int PRIMARY KEY ,
    code varchar(20),
    name varchar(100),
    status int default 0,   -- 0为有效，1为无效
    remark varchar(100)
);

create table sys_dict_data (
    id int PRIMARY KEY ,
    pid int ,
    seq int,
    code varchar(20),
    name varchar(100),
    ext1 varchar(20),       -- 扩展字段1
    ext2 varchar(20),       -- 扩展字段2
    ext3 varchar(200),      -- 扩展字段3
    status int default 0,   -- 0为有效，1为无效
    remark varchar(100)
);

insert into sys_dict_type(id,code,name) values (1,'payment_type','交易类型');
insert into sys_dict_data(id,pid,seq,code,name,remark) values (1,1,1,'incoming','incoming','入帐');
insert into sys_dict_data(id,pid,seq,code,name,remark) values (2,1,2,'outgoing','outgoing','出帐');
insert into sys_dict_data(id,pid,seq,code,name,remark) values (3,1,3,'frozen','frozen','冰结');

insert into sys_dict_type(id,code,name) values (2,'sys_currency','货币类型');
insert into sys_dict_data(id,pid,seq,code,name,remark) values (4,2,1,'usd','USD','美元');


create table account (
    id serial primary key ,
    code varchar(20),
    name varchar(100),
    balance int,          -- '以0.01元为单位'
    curr varchar(20),    -- 货币单位 sys_dict_data:4
    remark varchar(100)
);
-- balance int, frozen int,
insert into account(id,code,name,balance,curr) values (1,'a001','bob123',10000,'usd');
insert into account(id,code,name,balance,curr) values (2,'a002','alice456',1,'usd');

create table account_balance(
    id serial primary key,
    code varchar(20),        -- account.code
    curr varchar(20),        -- 货币单位 sys_dict_type:2
    balance int,              -- '以0.01元为单位'
    frozen int                -- '以0.01元为单位'
);
insert into account_balance(id,code,curr,balance,frozen) values (1,'a001','usd',10000,0);
insert into account_balance(id,code,curr,balance,frozen) values (2,'a002','usd',1,0);

create table payment(
    id serial primary key,
    account varchar(20),      -- account.code
    ptype varchar(20),      -- sys_dict_type:1
    curr varchar(20),      -- sys_dict_type:2
    balance int ,           -- '以0.01元为单位'
    frozen int,             -- '以0.01元为单位'
    token varchar(64),     -- 当前操作的token
    remark varchar(100),
    create_at int,          -- timestamp
    create_by varchar(20)   -- 操作员
);
create table payment_process(
    id serial primary key,
    account varchar(20),      -- account.code
    ptype varchar(20),      -- sys_dict_type:1
    curr varchar(20),      -- sys_dict_type:2
    balance int ,           -- '以0.01元为单位'
    frozen int,             -- '以0.01元为单位'
    token varchar(64),     -- 当前操作的token
    seq int,                -- 当前操作的次序
    remark varchar(100),
    create_at int,          -- timestamp
    create_by varchar(20)   -- 操作员
)
```