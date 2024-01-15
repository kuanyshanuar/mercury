ALTER TABLE flat_plans ALTER COLUMN images SET DEFAULT '{}';

UPDATE flat_plans
    SET images = '{}'
WHERE images IS NULL;