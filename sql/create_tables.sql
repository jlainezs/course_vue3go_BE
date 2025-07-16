create table public.users
(
    id         serial
        primary key,
    created_at timestamp default now(),
    email      varchar(255),
    first_name varchar(255),
    last_name  varchar(255),
    password   varchar(64),
    updated_at timestamp default now()
);

alter table public.users
    owner to postgres;
