create extension pgcrypto;

create table daryl (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(50) not null,
  password varchar(100),
);

create unique index daryl_name_index on daryl (name);
create index daryl_password_index on daryl (password);

create table habit (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  daryl_id UUID references daryl,
  title varchar(100) not null,
  duration varchar(100) not null,
  cron varchar(100) not null,
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
