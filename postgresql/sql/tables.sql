DROP TABLE IF EXISTS bucket;
DROP TABLE IF EXISTS object;

CREATE TABLE bucket (
    id TEXT NOT NULL UNIQUE PRIMARY KEY
);

CREATE TABLE object (
    id TEXT NOT NULL,
    content TEXT NOT NULL,
    bucket_id TEXT NOT NULL,
    PRIMARY KEY (id, bucket_id),
    FOREIGN KEY (bucket_id) REFERENCES bucket(id)
);