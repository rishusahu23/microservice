-- Snapshot for database: test_database
CREATE TABLE schema_migrations
(
version bigint,
dirty boolean
);


CREATE TABLE users
(
id uuid,
username character varying,
email character varying,
password_hash text,
created_at timestamp without time zone,
updated_at timestamp without time zone
);


