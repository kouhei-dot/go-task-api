create table if not exists Tasks (
    id SERIAL NOT NULL,
    uuid TEXT NOT NULL,
    label TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
