create table if not exists rooms (
	id bigserial primary key not null,
	room_name varchar(255) not null,
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now()
);