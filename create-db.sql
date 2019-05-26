-- psql postgres -f create-db.sql

DROP DATABASE IF EXISTS shorty;

DROP ROLE IF EXISTS shortyapp;

CREATE ROLE shortyapp LOGIN CREATEDB PASSWORD 'devpassword';

CREATE DATABASE shorty
  WITH OWNER = shortyapp
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       --LC_COLLATE = 'de_DE.UTF-8'
       --LC_CTYPE = 'de_DE.UTF-8'
       CONNECTION LIMIT = -1;
GRANT ALL ON DATABASE shorty TO shortyapp;
REVOKE ALL ON DATABASE shorty FROM public;
