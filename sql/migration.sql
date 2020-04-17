create table dbo.users
(
    id       int primary key identity (1, 1),
    name     varchar(50) not null,
    password varchar(50) not null
);

insert into dbo.users (name, password)
values ('admin', 'admin');

create table dbo.transactions
(
    id      int primary key identity (1,1),
    status  int not null,
    user_id int not null,
    foreign key (user_id) references dbo.users (id)
);

create table dbo.refunds
(
    id      int primary key identity (1,1),
    status  int not null,
    user_id int not null,
    foreign key (user_id) references dbo.users (id)
)

alter table dbo.refunds
    add amount money not null default 0;
alter table dbo.transactions
    add amount money not null default 0;

---

