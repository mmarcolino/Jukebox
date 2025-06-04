-- migrate:up
CREATE TABLE public.tracks(
    id VARCHAR (26) PRIMARY KEY,
    title VARCHAR (128) NOT NULL,
    artist VARCHAR (128) NOT NULL,
    album VARCHAR (128),
    duration INT NOT NULL,
    genre VARCHAR (64)
);

-- migrate:down
DROP TABLE public.tracks;
