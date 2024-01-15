CREATE TABLE IF NOT EXISTS flat_plans (
  id              BIGSERIAL NOT NULL PRIMARY KEY,
  residence_id    BIGINT REFERENCES residences (id),
  number_of_rooms INT NOT NULL,
  area            FLOAT NOT NULL,
  price           INT NOT NULL,
  images          VARCHAR[],
  created_at      BIGINT,
  updated_at      BIGINT,
  deleted_at      BIGINT
);