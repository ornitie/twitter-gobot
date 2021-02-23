CREATE TABLE IF NOT EXISTS Tweet (
    id SERIAL,
    tweet_id varchar(255),
    tweet_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id varchar(255),
    text text
);