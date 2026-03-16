CREATE TABLE comments (
    id BLOB PRIMARY KEY,
    site TEXT NOT NULL,
    resource TEXT NOT NULL,
    author TEXT NOT NULL,
    text TEXT NOT NULL,
    created_at DATETIME DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'))
);
