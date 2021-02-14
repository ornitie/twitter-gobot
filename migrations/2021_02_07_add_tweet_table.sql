CREATE TABLE Tweet (
    ID varchar(255),
    TweetDate TIMESTAMP,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    Text text(255)
);