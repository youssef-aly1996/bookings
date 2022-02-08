create table if not exists users (
	id bigserial primary key not null,
	first_name varchar(50) not null,
	last_name varchar(50) not null,
	email varchar(255) unique not null,
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
	access_level integer not null
);