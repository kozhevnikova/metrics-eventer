drop table device_metrics
drop table devices
drop table device_alerts
drop table users


create table users(
	id int primary key,
	name varchar(255) not null,
	email varchar(255) not null
)

create table devices(
	id int primary key,
	name varchar(255) not null,
	user_id int not null references users(id) on delete cascade
)

create table device_metrics(
	id serial primary key,
	device_id int not null references devices(id) on delete cascade,
	metric_1 int,
	metric_2 int,
	metric_3 int,
	metric_4 int,
	metric_5 int,
	local_time timestamp,
	server_time timestamp default now()
)
	
create table device_alerts(
	id serial primary key,
	device_id int,
	message text
)

insert into users(id,name,email) values
(1,'',''),
(2,'','');

do $$
declare userid integer := (select count(*) from users);
begin
	insert into devices(id,name,user_id) values
	(generate_series(1,1000),MD5(random()::text), generate_series(1,userid));
end $$;
