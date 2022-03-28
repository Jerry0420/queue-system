CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ----------------------------

CREATE TABLE IF NOT EXISTS stores(
   id serial PRIMARY KEY,
   email varchar(512) NOT NULL UNIQUE,
   password varchar(64) NOT NULL,
   name varchar(64) NOT NULL,
   timezone varchar(64) NOT NULL,
   description text DEFAULT '',
   created_at timestamp NOT NULL DEFAULT clock_timestamp()
);

-- ----------------------------

CREATE TYPE store_session_status AS ENUM ('normal', 'scanned', 'used');

CREATE TABLE IF NOT EXISTS store_sessions(
   id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
   store_id integer REFERENCES stores ON DELETE CASCADE,
   status store_session_status DEFAULT 'normal'
);

-- ----------------------------

CREATE TYPE sign_key_types AS ENUM ('normal', 'password', 'refresh', 'session');

CREATE TABLE IF NOT EXISTS sign_keys(
   id serial PRIMARY KEY,
   store_id integer REFERENCES stores ON DELETE CASCADE,
   sign_key varchar(64) NOT NULL,
   type sign_key_types NOT NULL
);

-- ----------------------------

CREATE TABLE IF NOT EXISTS queues(
   id serial PRIMARY KEY,
   name varchar(64) NOT NULL,
   store_id integer REFERENCES stores ON DELETE CASCADE
);

-- ----------------------------

CREATE TYPE customer_status AS ENUM ('waiting', 'processing', 'done', 'delete');

CREATE TABLE IF NOT EXISTS customers(
   id serial PRIMARY KEY,
   name varchar(64) NOT NULL,
   phone varchar(30) NOT NULL,
   queue_id integer REFERENCES queues ON DELETE CASCADE,
   status customer_status DEFAULT 'waiting',
   created_at timestamp NOT NULL DEFAULT clock_timestamp()
);

-- ----------------------------