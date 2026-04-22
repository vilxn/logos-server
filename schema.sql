CREATE TABLE users (
       id INTEGER PRIMARY KEY AUTOINCREMENT,

       email TEXT NOT NULL UNIQUE,
       password_hash TEXT NOT NULL,

       first_name TEXT NOT NULL,
       last_name TEXT NOT NULL,

       role TEXT NOT NULL CHECK(role IN ('parent', 'specialist', 'admin')),
       created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

