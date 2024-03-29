create extension pgcrypto;

create table daryl (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email varchar(50) not null,
  name varchar(50) not null,
  password varchar(100)
);

create unique index daryl_email_index on daryl (email);
create index daryl_password_index on daryl (password);

create table habit (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  daryl_id UUID references daryl,
  title varchar(100) not null,
  duration varchar(100) not null
);

create table habit_trigger (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  habit_id UUID references habit,
  daryl_id UUID references daryl,
  name varchar(100) not null,
  engine varchar(100) not null,
  params jsonb not null default '{}'::jsonb
);

create table message (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  daryl_id UUID references daryl,
  habit_id UUID references habit,
  text varchar not null,
  at timestamp with time zone,
  attrs jsonb not null default '{}'::jsonb
);

create index message_attrs_index on message using gin (attrs);
create index message_at_index on message (at);
