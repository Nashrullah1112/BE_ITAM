CREATE TABLE IF NOT EXISTS devices (
  id SERIAL PRIMARY KEY,
  recipient_location TEXT NULL,
  receipt_time DATE NULL,
  receipt_proof TEXT NULL,
  asset_type TEXT NULL,
  asset_activation_time DATE NULL,
  asset_inspection_result TEXT NULL,
  serial_number TEXT NULL,
  model TEXT NULL,
  warranty_start_time DATE NULL,
  warranty_card_number TEXT NULL,
  processor TEXT NULL,
  ram_capacity TEXT NULL,
  rom_capacity TEXT NULL,
  ram_type TEXT NULL,
  storage_type TEXT NULL,
  asset_status TEXT NULL,
  asset_value TEXT NULL,
  depreciation_value TEXT NULL,
  usage_period INT NULL,
  asset_out_time DATE NULL,
  asset_condition_on_exit TEXT NULL,
  purchase_receipt TEXT NULL,
  asset_id INT NULL,
  division_id INT NULL,
  user_id INT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (asset_id) REFERENCES assets (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (division_id) REFERENCES divisions (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);