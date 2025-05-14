-- Drop all tables in the correct order to avoid foreign key constraint issues

-- Drop notifications table
DROP TABLE IF EXISTS notifications;

-- Drop standouts table and constraints
DROP TABLE IF EXISTS standouts;

-- Drop messages table and constraints
DROP TABLE IF EXISTS messages;

-- Drop matches table and constraints
DROP TABLE IF EXISTS matches;

-- Drop swipes table and constraints
DROP TABLE IF EXISTS swipes;

-- Drop preferences table and constraints
DROP TABLE IF EXISTS preferences;

-- Drop profile_prompts table and constraints
DROP TABLE IF EXISTS profile_prompts;

-- Drop prompts table
DROP TABLE IF EXISTS prompts;

-- Drop photos table
DROP TABLE IF EXISTS photos;

-- Keep profiles and users tables
-- If you want to drop these as well, uncomment these:
-- DROP TABLE IF EXISTS profiles;
-- DROP TABLE IF EXISTS users;

-- Remove vices column from profiles table
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'profiles' AND column_name = 'vices') THEN
        ALTER TABLE profiles DROP COLUMN vices;
    END IF;
END
$$; 