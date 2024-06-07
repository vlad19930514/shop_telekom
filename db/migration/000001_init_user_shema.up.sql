CREATE TABLE users (
  id serial primary key,
  email varchar NOT NULL,
  password varchar NOT NULL,
  username varchar NOT NULL DEFAULT 'user',
  is_email_verified bool NOT NULL DEFAULT false
);
