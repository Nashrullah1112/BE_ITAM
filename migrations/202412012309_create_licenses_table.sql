CREATE TABLE IF NOT EXISTS licenses (
  id SERIAL PRIMARY KEY,
  purchase_date DATE NULL,
  installed_device_sn TEXT NULL,
  activation_date DATE NULL,
  expiration_date DATE NULL,
  asset_ownership_type TEXT NULL,
  license_category TEXT NULL,
  license_version TEXT NULL,
  max_application_users INT NULL,
  max_device_licenses INT NULL,
  license_type TEXT NULL,
  asset_id INT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (asset_id) REFERENCES assets (id) ON DELETE CASCADE ON UPDATE CASCADE
);