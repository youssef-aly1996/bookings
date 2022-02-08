create table if not exists room_restrictions (
	id bigserial primary key not null,
    start_date date not null,
	end_date date not null,
    room_id bigint not null,
    reservation_id bigint not null,
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
    foreign key (room_id) references rooms(id)  on delete cascade,
    foreign key (reservation_id) references reservations(id)  on delete cascade
);