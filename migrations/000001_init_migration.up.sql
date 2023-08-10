create table articles (
    id serial,
    author varchar(255) null,
    title varchar(255) null,
    body text null,
    created_at timestamp default now(),
    updated_at timestamp null,
    deleted_at timestamp null
)