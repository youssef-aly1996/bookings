create table if not exists restrictions (
    id serial primary key not null,
    restriction_name varchar(255),
    created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now()
)