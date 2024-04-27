CREATE TABLE public.accounts (
    id bigserial PRIMARY KEY,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NULL,
    deleted_at timestamp with time zone NULL,
    user_id bigint NOT NULL,
    exchange_id bigint NOT NULL,
    name text NOT NULL,
    capital text NOT NULL,
    locked_capital text NOT NULL
);
CREATE INDEX idx_accounts_del ON public.accounts USING HASH (deleted_at);
CREATE INDEX idx_accounts_user_id ON public.accounts USING HASH (user_id);
CREATE INDEX idx_accounts_exchange_id ON public.accounts USING HASH (exchange_id);