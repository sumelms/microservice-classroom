BEGIN;

CREATE TABLE classroom_lessons
(
    id                      bigserial       CONSTRAINT classroom_lessons_pk PRIMARY KEY,
    uuid                    uuid            DEFAULT uuid_generate_v4() NOT NULL,
    classroom_id            uuid,
    syllabus_lessons_id     uuid            NOT NULL,
    starts_at               timestamp       DEFAULT now() NOT NULL,
    ends_at                 timestamp
);

CREATE UNIQUE INDEX classroom_lessons_id_uindex
    ON classroom_lessons (id);
CREATE UNIQUE INDEX classroom_lessons_uuid_uindex
    ON classroom_lessons (uuid);

COMMIT;
