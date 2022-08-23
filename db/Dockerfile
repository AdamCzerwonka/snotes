FROM postgres:latest

ADD sql/create_db.sql /docker-entrypoint-initdb.d
ADD sql/create_tables.sql /docker-entrypoint-initdb.d

