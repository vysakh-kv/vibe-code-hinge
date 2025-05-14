-- Drop tables in reverse order of creation to respect foreign key constraints
DROP TABLE IF EXISTS user_preferences;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS swipes;
DROP TABLE IF EXISTS matches;
DROP TABLE IF EXISTS photos;
DROP TABLE IF EXISTS profiles;
DROP TABLE IF EXISTS users; 