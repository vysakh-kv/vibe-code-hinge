-- Create or update users table if not exists
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create profiles table if not exists
CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY,
    user_id UUID UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    bio TEXT,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(50) NOT NULL,
    location VARCHAR(255),
    occupation VARCHAR(255),
    vices JSONB DEFAULT '{}'::JSONB,
    preferences JSONB DEFAULT '{}'::JSONB,
    onboarding_completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT profiles_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Add photos table
CREATE TABLE IF NOT EXISTS photos (
    id BIGSERIAL PRIMARY KEY,
    profile_id UUID NOT NULL,
    url VARCHAR(255) NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT photos_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE
);

-- Create prompts table
CREATE TABLE IF NOT EXISTS prompts (
    id BIGSERIAL PRIMARY KEY,
    text VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create profile_prompts table
CREATE TABLE IF NOT EXISTS profile_prompts (
    id BIGSERIAL PRIMARY KEY,
    profile_id UUID NOT NULL,
    prompt_id BIGINT NOT NULL,
    answer TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(profile_id, prompt_id),
    CONSTRAINT profile_prompts_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE,
    CONSTRAINT profile_prompts_prompt_id_fkey FOREIGN KEY (prompt_id) REFERENCES prompts(id)
);

-- Create preferences table
CREATE TABLE IF NOT EXISTS preferences (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID UNIQUE NOT NULL,
    preferred_gender VARCHAR(50),
    min_age INT CHECK (min_age >= 18),
    max_age INT CHECK (max_age <= 100),
    max_distance INT,
    preferences JSONB DEFAULT '{}'::JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT preferences_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create swipes table
CREATE TABLE IF NOT EXISTS swipes (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    profile_id UUID NOT NULL,
    is_like BOOLEAN NOT NULL,
    message TEXT,
    is_rose BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, profile_id),
    CONSTRAINT swipes_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT swipes_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE
);

-- Create matches table
CREATE TABLE IF NOT EXISTS matches (
    id BIGSERIAL PRIMARY KEY,
    user1_id UUID NOT NULL,
    user2_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_message_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    user1_last_read TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    user2_last_read TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user1_id, user2_id),
    CONSTRAINT matches_user1_id_fkey FOREIGN KEY (user1_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT matches_user2_id_fkey FOREIGN KEY (user2_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create messages table
CREATE TABLE IF NOT EXISTS messages (
    id BIGSERIAL PRIMARY KEY,
    match_id BIGINT NOT NULL,
    sender_id UUID NOT NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT messages_match_id_fkey FOREIGN KEY (match_id) REFERENCES matches(id) ON DELETE CASCADE,
    CONSTRAINT messages_sender_id_fkey FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create standouts table
CREATE TABLE IF NOT EXISTS standouts (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    profile_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    expires_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() + INTERVAL '1 day',
    is_active BOOLEAN DEFAULT TRUE,
    UNIQUE(user_id, profile_id),
    CONSTRAINT standouts_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT standouts_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE
);

-- Create notifications table
CREATE TABLE IF NOT EXISTS notifications (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    type VARCHAR(50) NOT NULL,
    target_id BIGINT,
    message TEXT,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT notifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Insert default prompts
INSERT INTO prompts (text) VALUES 
('Two truths and a lie...'),
('I''m looking for...'),
('You should leave a comment if...'),
('The key to my heart is...'),
('I take pride in...'),
('My most irrational fear...'),
('A life goal of mine...'),
('My simple pleasures...'),
('I get along best with people who...'),
('I''m convinced that...'),
('My most controversial opinion...'),
('I want someone who...'),
('We''ll get along if...'),
('My typical Sunday...'),
('Dating me is like...')
ON CONFLICT DO NOTHING; 