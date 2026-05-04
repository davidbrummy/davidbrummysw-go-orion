create table if not exists users (
	id bigserial primary key,
	uuid text not null unique,
	user_name text not null
);

insert into users (uuid, user_name)
values ('test-uuid', 'Test Name')
on conflict (uuid) do nothing;
