-- SEQUENCE: public.user_id_seq

-- DROP SEQUENCE public.user_id_seq;

CREATE SEQUENCE public.user_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.user_id_seq
    OWNER TO postgres;

-- Table: public."user"

-- DROP TABLE public."user";

CREATE TABLE public."user"
(
    id integer NOT NULL DEFAULT nextval('user_id_seq'::regclass),
    CONSTRAINT "User_pkey" PRIMARY KEY (id)
)

    TABLESPACE pg_default;

ALTER TABLE public."user"
    OWNER to postgres;

-- SEQUENCE: public.transaction_id_seq

-- DROP SEQUENCE public.transaction_id_seq;

CREATE SEQUENCE public.transaction_id_seq
    INCREMENT 1
    START 66
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.transaction_id_seq
    OWNER TO postgres;

-- Table: public.transaction

-- DROP TABLE public.transaction;

CREATE TABLE public.transaction
(
    id integer NOT NULL DEFAULT nextval('transaction_id_seq'::regclass),
    user_id integer NOT NULL,
    amount numeric(16,2) NOT NULL,
    type smallint NOT NULL,
    service_id integer,
    order_id integer,
    date timestamp without time zone NOT NULL,
    CONSTRAINT transaction_pkey PRIMARY KEY (id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

    TABLESPACE pg_default;

ALTER TABLE public.transaction
    OWNER to postgres;
