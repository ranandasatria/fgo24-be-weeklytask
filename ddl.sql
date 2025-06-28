-- Active: 1750753580331@@127.0.0.1@5433@ewallet
CREATE DATABASE ewallet;

CREATE TABLE users (
  id_user SERIAL PRIMARY KEY,
  email VARCHAR(100) UNIQUE NOT NULL,
  password TEXT NOT NULL,
  pin VARCHAR(6) NOT NULL,
  username VARCHAR(50) NOT NULL,
  phone VARCHAR(20),
  profile_picture VARCHAR(250)
);

ALTER TABLE users
ALTER COLUMN password TYPE VARCHAR(250);

CREATE TABLE sessions (
  id_session SERIAL PRIMARY KEY,
  id_user INT REFERENCES users(id_user),
  issued_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE wallets (
  id_wallet SERIAL PRIMARY KEY,
  id_user INT UNIQUE REFERENCES users(id_user),
  balance DECIMAL(12, 2) DEFAULT 0 CHECK (balance >= 0)
);

CREATE TABLE payment_methods (
  id_payment_method SERIAL PRIMARY KEY,
  payment_method VARCHAR(50) NOT NULL
);

CREATE TABLE topups (
  id_topup SERIAL PRIMARY KEY,
  id_wallet INT REFERENCES wallets(id_wallet),
  amount DECIMAL(12, 2) NOT NULL CHECK (amount > 0),
  id_payment_method INT REFERENCES payment_methods(id_payment_method),
  admin_fee DECIMAL(12, 2) DEFAULT 0,
  tax DECIMAL(12, 2) DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transfers (
  id_transfer SERIAL PRIMARY KEY,
  id_sender_wallet INT REFERENCES wallets(id_wallet),
  id_receiver_wallet INT REFERENCES wallets(id_wallet),
  amount DECIMAL(12, 2) NOT NULL CHECK (amount > 0),
  notes TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
