CREATE TABLE words (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    japanese TEXT,
    romaji TEXT,
    english TEXT,
    parts JSON
);

CREATE TABLE word_groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    word_id INTEGER,
    group_id INTEGER
);

CREATE TABLE groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT
);

CREATE TABLE study_sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER,
    created_at DATETIME,
    study_activity_id INTEGER,
    activity_name TEXT,
    group_name TEXT
);

CREATE TABLE study_activities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    study_session_id INTEGER,
    group_id INTEGER,
    created_at DATETIME
);

CREATE TABLE word_review_items (
    word_id INTEGER,
    study_session_id INTEGER,
    correct BOOLEAN,
    created_at DATETIME
);