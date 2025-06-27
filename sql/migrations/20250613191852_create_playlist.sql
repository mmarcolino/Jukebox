-- migrate:up
CREATE TABLE public.playlist(
    id VARCHAR (26) PRIMARY KEY,
    name VARCHAR (128) NOT NULL,
    tracks VARCHAR(26)[] NOT NULL
);

-- migrate:down
DROP TABLE public.playlist;