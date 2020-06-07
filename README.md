# Table Creation Query: 
    create table if not exists products (
        `id` bigint not null auto_increment primary key,
        `created_at` datetime(6) not null,
        `modified_at` datetime(6) not null,
        `name` VARCHAR(255) not null,
        `price` int not null,
        `deleted_at` datetime(6) default null,
        `is_deletd` tinyint(1) not null default 0
    )