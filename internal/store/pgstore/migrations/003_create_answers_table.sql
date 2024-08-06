-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS answers
(
    "id"            uuid            PRIMARY KEY DEFAULT gen_random_uuid(),
    "message_id"    uuid            NOT NULL,
    "answer"        VARCHAR(255)    NOT NULL,

    FOREIGN KEY (message_id) REFERENCES messages(id)
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
