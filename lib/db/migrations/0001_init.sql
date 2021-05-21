create table if not exists "USER"
(
    id   serial primary key,
    name varchar not null
);

create table if not exists widget
(
    id       serial primary key,
    type     varchar                       not null,
    owner_id serial references "USER" (id) not null
);
