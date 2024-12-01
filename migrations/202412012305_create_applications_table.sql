CREATE TABLE IF NOT EXISTS applications (
  id SERIAL PRIMARY KEY,
  application_name TEXT NULL,
  creation_date DATE NULL,
  acceptance_date DATE NULL,
  storage_server_location TEXT NULL,
  application_type TEXT NULL,
  application_link TEXT NULL,
  application_certification TEXT NULL,
  activation_date DATE NULL,
  expiration_date DATE NULL,
  asset_id INT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (asset_id) REFERENCES assets (id) ON DELETE CASCADE ON UPDATE CASCADE
);