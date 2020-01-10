CREATE TABLE IF NOT EXISTS cities(
   id SERIAL PRIMARY KEY,
   name VARCHAR (300) UNIQUE NOT NULL,
   latitude float,
   longitude float,
   is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS temperatures(
   id SERIAL PRIMARY KEY,
   city_id     integer REFERENCES cities (id),
   max integer,
   min integer,
   timestamp integer NOT NULL
);

CREATE TABLE IF NOT EXISTS webhooks(
   id SERIAL PRIMARY KEY,
   city_id     integer REFERENCES cities (id),
   callback_url TEXT,
   is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

