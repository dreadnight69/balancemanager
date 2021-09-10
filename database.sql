CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE CACHE 1;

CREATE SEQUENCE public.transactions_transaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE CACHE 1;

CREATE TABLE public.users
(
    user_id bigint NOT NULL DEFAULT nextval('public.users_user_id_seq'::regclass),
    balance bigint NOT NULL CHECK (balance >= 0),

    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE TABLE public.transactions
(
    transaction_id    bigint      NOT NULL DEFAULT nextval('public.transactions_transaction_id_seq'::regclass),
    initiator_id      bigint,
    recipient_id      bigint,
    amount            bigint      NOT NULL,
    operation_type_id int CHECK (operation_type_id >= 1 AND operation_type_id <= 3),
    description       text,
    date              timestamptz NOT NULL,

    CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id),
    CONSTRAINT transactions_initiator_id_fkey FOREIGN KEY (initiator_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT transactions_recipient_id_fkey FOREIGN KEY (recipient_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);