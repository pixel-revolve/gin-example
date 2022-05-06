create database if not exists `gin`;

use `gin`;

drop table if exists `user`;

create table if not exists `user`(
    `id` bigint primary key auto_increment,
    `uuid` binary(16),
    `username` varchar(30),
    `password` varchar(30),
    `nickname` varchar(30),
    `phone` varchar(30),
    `email` varchar(40),
    `create_at` datetime,
    `update_at` datetime,
    `deleted_at` datetime

)engine=Innodb