DROP TABLE IF EXISTS examplegifts;
CREATE TABLE IF NOT EXISTS examplegifts (
                              gift_id integer PRIMARY KEY,
                              name varchar NOT NULL,
                              price integer NOT NULL
);