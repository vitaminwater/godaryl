create extension pgcrypto;

create table daryl (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(50) not null,
  password varchar(100),
  attrs jsonb not null default '{}'::jsonb
);

create unique index daryl_name_index on daryl (name);
create index daryl_password_index on daryl (password);
create index daryl_attrs_index on daryl using gin (attrs);

create table habit (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  daryl_id UUID references daryl,
  title varchar(100) not null,
  duration varchar(100) not null,
  deadline date,
  cron varchar(100) not null,
  attrs jsonb not null default '{}'::jsonb
);

create index habit_attrs_index on habit using gin (attrs);

create table message (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  daryl_id UUID references daryl,
  habit_id UUID references habit,
  text varchar not null,
  at timestamp with time zone,
  attrs jsonb not null default '{}'::jsonb
);

create index message_attrs_index on message using gin (attrs);

create table eventslog (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  daryl_id UUID references daryl,
  eventtype varchar(10) not null,
  attrs jsonb not null default '{}'::jsonb
);

create index eventslog_eventtype_index on eventslog (eventtype);
create index eventslog_attrs_index on eventslog using gin (attrs);
