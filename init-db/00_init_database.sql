CREATE USER chat_app_user WITH ENCRYPTED PASSWORD 'test1234';

CREATE DATABASE chat_db;
GRANT ALL PRIVILEGES ON DATABASE chat_db TO chat_app_user;

\c chat_db;

GRANT USAGE ON SCHEMA public TO chat_app_user;
GRANT CREATE ON SCHEMA public TO chat_app_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO chat_app_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT ON SEQUENCES TO chat_app_user;


-- 테이블 생성 (이미 존재하는지 확인 후 생성)
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS rooms (
    room_id SERIAL PRIMARY KEY,
    room_name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS room_members (
    room_id INTEGER REFERENCES rooms(room_id),
    member_id INTEGER REFERENCES users(user_id),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS messages (
    message_id SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES rooms(room_id),
    sender_id INTEGER REFERENCES users(user_id),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);