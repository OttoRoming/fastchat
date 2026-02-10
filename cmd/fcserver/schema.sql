create table account (
    id integer primary key not null,
    username text not null,
    password text not null
);

create table chat (
    id integer primary key not null,
    from_id integer not null,
    to_id integer not null,
    content text not null
    foreign key (from_id) references account(id)
    foreign key (to_id) references account(id)
);

create index idx_chat_from_id on chat(from_id);
create index idx_chat_to_id on chat(to_id);
