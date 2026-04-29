PRAGMA foreign_keys = ON;

-- =========================
-- USERS (родители, специалисты, админ)
-- =========================
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL CHECK(role IN ('parent', 'specialist', 'admin')),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    is_approved INTEGER DEFAULT 0,
    is_banned INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE parents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL UNIQUE,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE children (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    birth_date DATE,
    notes TEXT
);

CREATE TABLE child_parents (
    child_id INTEGER NOT NULL,
    parent_id INTEGER NOT NULL,
    FOREIGN KEY(child_id) REFERENCES children(id) ON DELETE CASCADE,
    FOREIGN KEY(parent_id) REFERENCES parents(id) ON DELETE CASCADE
);

-- =========================
-- CHILD ↔ SPECIALIST (многие ко многим)
-- =========================
CREATE TABLE child_specialists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    child_id INTEGER NOT NULL,
    specialist_id INTEGER NOT NULL,

    UNIQUE(child_id, specialist_id),

    FOREIGN KEY(child_id) REFERENCES children(id) ON DELETE CASCADE,
    FOREIGN KEY(specialist_id) REFERENCES users(id) ON DELETE CASCADE
);

-- =========================
-- APPOINTMENTS (записи на занятия)
-- =========================
CREATE TABLE appointments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    child_id INTEGER NOT NULL,
    specialist_id INTEGER NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    status TEXT NOT NULL CHECK(status IN ('pending', 'confirmed', 'cancelled')),
    created_by INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(child_id) REFERENCES children(id) ON DELETE CASCADE,
    FOREIGN KEY(specialist_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(created_by) REFERENCES users(id)
);

-- =========================
-- REPORTS (отчеты специалистов)
-- =========================
CREATE TABLE reports (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    child_id INTEGER NOT NULL,
    specialist_id INTEGER NOT NULL,
    appointment_id INTEGER,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(child_id) REFERENCES children(id) ON DELETE CASCADE,
    FOREIGN KEY(specialist_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(appointment_id) REFERENCES appointments(id) ON DELETE SET NULL
);

-- =========================
-- PROGRESS (метрики прогресса)
-- =========================
CREATE TABLE progress (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    child_id INTEGER NOT NULL,
    metric_name TEXT NOT NULL,
    value REAL NOT NULL,
    recorded_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(child_id) REFERENCES children(id) ON DELETE CASCADE
);

-- =========================
-- RECOMMENDATIONS (рекомендации)
-- =========================
CREATE TABLE recommendations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    child_id INTEGER NOT NULL,
    specialist_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    type TEXT NOT NULL CHECK(type IN ('child', 'parent')),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(child_id) REFERENCES children(id) ON DELETE CASCADE,
    FOREIGN KEY(specialist_id) REFERENCES users(id) ON DELETE CASCADE
);

-- =========================
-- COURSES (курсы для родителей)
-- =========================
CREATE TABLE courses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    created_by INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(created_by) REFERENCES users(id) ON DELETE CASCADE
);

-- =========================
-- COURSE ENROLLMENTS (запись на курсы)
-- =========================
CREATE TABLE course_enrollments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    course_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    enrolled_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(course_id, user_id),

    FOREIGN KEY(course_id) REFERENCES courses(id) ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- =========================
-- LOGS (логи действий)
-- =========================
CREATE TABLE logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    action TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL
);

-- =========================
-- INDEXES
-- =========================
CREATE INDEX idx_child_specialists_child ON child_specialists(child_id);
CREATE INDEX idx_child_specialists_specialist ON child_specialists(specialist_id);

CREATE INDEX idx_appointments_child ON appointments(child_id);
CREATE INDEX idx_appointments_specialist ON appointments(specialist_id);
CREATE INDEX idx_appointments_time ON appointments(start_time, end_time);

CREATE INDEX idx_reports_child ON reports(child_id);
CREATE INDEX idx_progress_child ON progress(child_id);

CREATE INDEX idx_recommendations_child ON recommendations(child_id);
CREATE INDEX idx_courses_creator ON courses(created_by);

