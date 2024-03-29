BEGIN;

CREATE TABLE classrooms
(
    id              bigserial       CONSTRAINT classrooms_pk PRIMARY KEY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL UNIQUE,
    name            varchar         NOT NULL,
    description     text,
    format          varchar         DEFAULT 'online',
    can_subscribe   bool            DEFAULT false,
    subject_id      uuid,
    course_id       uuid            NOT NULL,
    starts_at       timestamp       DEFAULT now() NOT NULL,
    ends_at         timestamp,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX classrooms_id_uindex
    ON classrooms (id);
CREATE UNIQUE INDEX classrooms_uuid_uindex
    ON classrooms (uuid);

COMMIT;
