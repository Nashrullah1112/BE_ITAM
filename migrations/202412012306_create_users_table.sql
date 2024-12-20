CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  nip TEXT NULL,
  name TEXT NULL,
  email TEXT UNIQUE NULL,
  password TEXT NULL,
  role TEXT NULL,
  join_date DATE NULL,
  division_id INT NULL,
  position_id INT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (division_id) REFERENCES divisions (id) ON DELETE SET NULL ON UPDATE CASCADE,
  FOREIGN KEY (position_id) REFERENCES positions (id) ON DELETE SET NULL ON UPDATE CASCADE
);