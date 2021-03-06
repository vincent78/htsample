drop table if exists sys_dict_type;
create table sys_dict_type
(
    id     serial PRIMARY KEY,
    code   varchar(20),
    name   varchar(100),
    status int default 0, -- 0为有效，1为无效
    remark varchar(100)
);
drop table if exists sys_dict_data;
create table sys_dict_data
(
    id     serial primary key,
    pid    int,
    seq    int,
    code   varchar(20),
    name   varchar(100),
    ext1   varchar(20),   -- extend1
    ext2   varchar(20),   -- extend2
    ext3   varchar(200),  -- extend3
    status int default 0, --
    remark varchar(100)
);

insert into sys_dict_type(id, code, name)
values (1, 'payment_type', 'paymentType');
insert into sys_dict_data(id, pid, seq, code, name, remark)
values (1, 1, 1, 'incoming', 'incoming', 'incoming');
insert into sys_dict_data(id, pid, seq, code, name, remark)
values (2, 1, 2, 'outgoing', 'outgoing', 'outgoing');
insert into sys_dict_data(id, pid, seq, code, name, remark)
values (3, 1, 3, 'frozen', 'frozen', 'frozen');
insert into sys_dict_data(id, pid, seq, code, name, remark)
values (4, 1, 4, 'commit', 'commit', 'commit');

insert into sys_dict_type(id, code, name)
values (2, 'sys_currency', 'currencyName');
insert into sys_dict_data(id, pid, seq, code, name, remark)
values (5, 2, 1, 'usd', 'USD', 'usd');

drop table if exists account;
create table account
(
    id     serial primary key,
    code   varchar(20),
    name   varchar(100),
    remark varchar(100)
);
insert into account(id, code, name)
values (1, 'a001', 'bob123');
insert into account(id, code, name)
values (2, 'a002', 'alice456');

drop table if exists account_balance;
create table account_balance
(
    id      serial primary key,
    code    varchar(20), -- account.code
    curr    varchar(20), -- sys_dict_type:2
    balance bigint       -- base 0.01
);
insert into account_balance(id, code, curr, balance)
values (1, 'a001', 'usd', 10000);
insert into account_balance(id, code, curr, balance)
values (2, 'a002', 'usd', 1);

drop table if exists payment;
create table payment
(
    id        serial primary key,
    account   varchar(20), -- account.code
    ptype     varchar(20), -- sys_dict_type:1
    curr      varchar(20), -- sys_dict_type:2
    balance   bigint,      -- base 0.01
    frozen    bigint,      -- base 0.01
    token     varchar(64), -- token
    remark    varchar(100),
    create_at bigint,      -- timestamp
    create_by varchar(20)  --
);

drop table if exists payment_process;
create table payment_process
(
    id        serial primary key,
    account   varchar(20), -- account.code
    ptype     varchar(20), -- sys_dict_type:1
    curr      varchar(20), -- sys_dict_type:2
    num       bigint,      -- base 0.01
    token     varchar(64), -- token
    seq       int,         -- sequence
    remark    varchar(100),
    create_at bigint,      -- timestamp
    create_by varchar(20)  -- s
)