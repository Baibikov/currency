create table currency.pairs(
   currency_from varchar(50) not null,
   currency_to   varchar(50) not null,
   well          decimal   not null default 0,
   updated_at    timestamptz not null default now()
);

alter table currency.pairs
    add constraint currency_pairs_currency_from_currency_to_uniq_key unique (currency_from, currency_to)
;