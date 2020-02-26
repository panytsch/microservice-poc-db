create table master.users
(
    id       int primary key identity (1, 1),
    name     varchar(50) not null,
    password varchar(50) not null
);

insert into master.users (name, password)
values ('admin', 'admin');

create table master.transactions
(
    id int primary key identity (1,1),
    status int not null,
    user_id int not null,
    foreign key (user_id) references master.users (id)
)

