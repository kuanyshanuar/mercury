CREATE TABLE  leads_statuses (
    id   BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR
);

INSERT INTO leads_statuses (id, name)
    VALUES (1, 'Активный'),
           (2, 'Неактивный');

ALTER TABLE residence_leads
    ADD COLUMN IF NOT EXISTS id BIGSERIAL NOT NULL PRIMARY KEY;

ALTER TABLE residence_leads
    ADD COLUMN IF NOT EXISTS status_id BIGINT DEFAULT 1 REFERENCES leads_statuses(id);