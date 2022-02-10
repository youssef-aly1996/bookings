create table if not exists reservations (
	id bigserial primary key not null,
	first_name varchar(50) not null,
	last_name varchar(50) not null,
	email varchar(255) unique not null,
    phone varchar(50) not null,
    start_date date not null,
	end_date date not null,
    room_id bigint not null,
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
    foreign key (room_id) references rooms(id)  on delete cascade
);