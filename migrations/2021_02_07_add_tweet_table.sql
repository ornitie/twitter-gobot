CREATE TABLE IF NOT EXISTS Tweet (
    id SERIAL,
    tweet_id varchar(255),
    tweet_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    author_id varchar(255),
    text text
);