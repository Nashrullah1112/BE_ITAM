CREATE TABLE IF NOT EXISTS vendors (
  id SERIAL PRIMARY KEY,
  contact_person TEXT NULL,
  email TEXT NULL,
  contact_number TEXT NULL,
  location TEXT NULL,
  siup_number TEXT NULL,
  nib_number TEXT NULL,
  npwp_number TEXT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);