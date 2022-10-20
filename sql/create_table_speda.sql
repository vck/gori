create table speda_fleet (
    id int primary key auto_increment,
    is_available boolean default false,
    fleet_name varchar(50) default null,
    created_date timestamp default now()
)