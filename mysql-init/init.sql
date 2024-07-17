DROP DATABASE IF EXISTS test_db;

CREATE DATABASE test_db;

USE test_db;

CREATE TABLE events (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  location VARCHAR(255) NOT NULL,
  organization VARCHAR(255) NOT NULL,
  rating VARCHAR(10) NOT NULL,
  date DATETIME NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  capacity INT NOT NULL,
  price FLOAT NOT NULL,
  partner_id INT NOT NULL
);

CREATE TABLE spots (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  event_id VARCHAR(36) NOT NULL,
  name VARCHAR(10) NOT NULL,
  status VARCHAR(10) NOT NULL,
  ticket_id VARCHAR(36),
  FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE TABLE tickets (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  event_id VARCHAR(36) NOT NULL,
  spot_id VARCHAR(36) NOT NULL,
  ticket_kind VARCHAR(10) NOT NULL,
  price FLOAT NOT NULL,
  FOREIGN KEY (event_id) REFERENCES events(id),
  FOREIGN KEY (spot_id) REFERENCES spots(id)
);

INSERT INTO events (id, name, location, organization, rating, date, image_url, capacity, price, partner_id) VALUES
  ('10853e59-dc5b-4d7b-a028-01513ef50d76', 'Event 001 - Partner1', 'São Paulo, SP', 'Partner 1', 'L14', '2021-10-10 10:00:00', 'https://images.unsplash.com/photo-1470229722913-7c0e2dbbafd3', 10, 100, 1),
  ('e0352b32-7698-4805-b029-28302b3a911f', 'Event 002 - Partner1', 'Rio de Janeiro, RJ', 'Partner 1', 'L14', '2021-10-10 12:00:00', 'https://images.unsplash.com/photo-1459749411175-04bf5292ceea', 10, 200, 1),
  ('5b79831a-a9d3-4538-8fb5-569494bd17a5', 'Event 003 - Partner2', 'Belo Horizonte, MG', 'Partner 2', 'L12', '2024-10-10 10:00:00', 'https://images.unsplash.com/photo-1540039155733-5bb30b53aa14', 10, 400, 2),
  ('8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'Event 004 - Partner2', 'Uberlândia, MG', 'Partner 2', 'L16', '2024-10-10 12:00:00', 'https://images.unsplash.com/photo-1493225457124-a3eb161ffa5f', 10, 500, 2)
;

