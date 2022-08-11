BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE classroom_lessons
(
    id                      bigserial       CONSTRAINT classroom_lessons_pk PRIMARY KEY,
    classroom_id            uuid,
    syllabus_lessons_id     uuid            NOT NULL,
    starts_at               timestamp       DEFAULT now() NOT NULL,
    ends_at                 timestamp
);

CREATE UNIQUE INDEX classrooms_id_uindex
    ON classrooms (id);

COMMIT;
