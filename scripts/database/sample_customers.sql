CREATE TABLE IF NOT EXISTS public.customers
(
    id integer PRIMARY KEY,
    name text,
    email_address text
);

--TABLESPACE pg_default;

\COPY customers(id,name,email_address) FROM 'customers.csv' DELIMITER ',' CSV HEADER;