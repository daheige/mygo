create database test default charset utf8;
set names utf8;

use test;

create table `user_info` (
    `id` int(10) unsigned not null primary key auto_increment,
    `username` varchar(64) not null default '',
    `depart` varchar(32) not null default '',
    `created` date null default null
);

create table `test` (
    `id` int(10) unsigned not null primary key auto_increment,
    `name` varchar(64) not null default ''
);

