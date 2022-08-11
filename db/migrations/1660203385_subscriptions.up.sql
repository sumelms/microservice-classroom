BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subscriptions
(
    id              bigserial       CONSTRAINT subscriptions_pk PRIMARY KEY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    user_id         uuid            NOT NULL,
    classroom_id    uuid            NOT NULL,
    role            varchar         NULL,
    expires_at      timestamp       NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX subscriptions_id_uindex
    ON subscriptions (id);

COMMIT;
