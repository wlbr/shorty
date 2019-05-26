--psql shorty -U shortyapp -f create-tables.sql

DROP TABLE IF EXISTS shortnames CASCADE;

CREATE TABLE shortnames (
  --id serial NOT NULL primary key,
 	localpart varchar(80) NOT NULL UNIQUE primary key,
 	longurl varchar(1024),
 	username varchar(30)
 	--userid int
 	);

ALTER TABLE shortnames
  OWNER TO shortyapp;

CREATE UNIQUE INDEX localpart_idx ON shortnames (localpart);
CREATE INDEX username_idx ON shortnames (username);



