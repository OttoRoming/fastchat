CREATE TABLE account (
    id TEXT PRIMARY KEY NOT NULL,
    token TEXT UNIQUE NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE chat (
    id TEXT PRIMARY KEY NOT NULL,
    from_id INTEGER NOT NULL,
    to_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY (from_id) REFERENCES account(id),
    FOREIGN KEY (to_id) REFERENCES account(id)
);

CREATE INDEX idx_chat_from_id ON chat(from_id);
CREATE INDEX idx_chat_to_id ON chat(to_id);
