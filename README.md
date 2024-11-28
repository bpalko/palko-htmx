# palko-htmx
HTMX sandbox

## Create DB Tables
```
touch gunbuild.db
sqlite3 gunbuild.db

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS components (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    brand TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS builds (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    barrel TEXT NOT NULL,
    grip TEXT NOT NULL,
    sight TEXT NOT NULL
);

```