create database daryl;

\c daryl;

create table primate (
  id serial,
  slug varchar(36) not null,
  username varchar(40) not null
);

create unique index primate_name_index on primate (slug);
