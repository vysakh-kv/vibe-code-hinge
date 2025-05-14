# Vibe Dating App API Curl Commands

Below are curl commands for all endpoints in the Vibe Dating App API. You can use these to test the API or import them into Postman.

## Base URL

```
BASE_URL="http://localhost:8082/api/v1"
```

## Health Check

```bash
# Health Check
curl -X GET "${BASE_URL}/health"
```

## Authentication

```bash
# Register a new user
curl -X POST "${BASE_URL}/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "Password123!"
  }'

# Login
curl -X POST "${BASE_URL}/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "Password123!"
  }'
```

## Profile Management

```bash
# Get profile by ID
curl -X GET "${BASE_URL}/profiles/{id}?user_id={user_id}"

# Update profile
curl -X PUT "${BASE_URL}/profiles/{id}?user_id={user_id}" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith",
    "bio": "Software engineer who loves hiking",
    "date_of_birth": "1990-01-01",
    "gender": "male",
    "location": "New York",
    "occupation": "Software Engineer",
    "photos": [
      "https://example.com/photo1.jpg",
      "https://example.com/photo2.jpg"
    ],
    "vices": {
      "drinking": "socially",
      "smoking": "never"
    }
  }'

# Get profile prompts
curl -X GET "${BASE_URL}/profiles/{id}/prompts?user_id={user_id}"

# Update profile prompt
curl -X PUT "${BASE_URL}/profiles/{id}/prompts?user_id={user_id}" \
  -H "Content-Type: application/json" \
  -d '{
    "prompt_id": "1",
    "response": "My response to this prompt"
  }'
```

## Preferences

```bash
# Get preferences
curl -X GET "${BASE_URL}/preferences?user_id={user_id}"

# Update preferences
curl -X PUT "${BASE_URL}/preferences?user_id={user_id}" \
  -H "Content-Type: application/json" \
  -d '{
    "preferred_gender": "female",
    "min_age": 21,
    "max_age": 35,
    "max_distance": 25,
    "preferences": {
      "height_min": 160,
      "height_max": 185,
      "relationship_goal": "casual"
    }
  }'
```

## Prompts

```bash
# Get default prompts
curl -X GET "${BASE_URL}/prompts"
```

## Feed and Discovery

```bash
# Get feed
curl -X GET "${BASE_URL}/feed?user_id={user_id}&limit=10&offset=0"

# Get standouts
curl -X GET "${BASE_URL}/standouts?user_id={user_id}&limit=5"

# Get discover profiles (legacy)
curl -X GET "${BASE_URL}/profiles/discover?user_id={user_id}&limit=10"
```

## Interaction (Swipes, Likes, etc.)

```bash
# Like a profile
curl -X POST "${BASE_URL}/profiles/{profile_id}/like?user_id={user_id}"

# Skip a profile
curl -X POST "${BASE_URL}/profiles/{profile_id}/skip?user_id={user_id}"

# Send a rose to a profile (premium feature)
curl -X POST "${BASE_URL}/profiles/{profile_id}/rose?user_id={user_id}"

# Create a swipe
curl -X POST "${BASE_URL}/swipes?user_id={user_id}" \
  -H "Content-Type: application/json" \
  -d '{
    "profile_id": "{profile_id}",
    "is_like": true,
    "message": "I noticed we both like hiking!",
    "is_rose": false
  }'
```

## Matches and Likes

```bash
# Get likes received
curl -X GET "${BASE_URL}/likes?user_id={user_id}"

# Get matches
curl -X GET "${BASE_URL}/matches?user_id={user_id}"
```

## Messaging

```bash
# Get messages for a match
curl -X GET "${BASE_URL}/matches/{match_id}/messages?user_id={user_id}&limit=20&offset=0"

# Send a message
curl -X POST "${BASE_URL}/matches/{match_id}/messages?user_id={user_id}" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "Hey, how are you doing?"
  }'
```

## Real-time Events (SSE)

```bash
# Connect to message events stream
curl -X GET "${BASE_URL}/events/messages?user_id={user_id}" \
  -H "Accept: text/event-stream" \
  -H "Cache-Control: no-cache" \
  -N

# Connect to notification events stream
curl -X GET "${BASE_URL}/events/notifications?user_id={user_id}" \
  -H "Accept: text/event-stream" \
  -H "Cache-Control: no-cache" \
  -N
```

## Importing to Postman

1. Copy any of these curl commands.
2. Open Postman.
3. Click "Import" in the top left.
4. Choose the "Raw text" tab.
5. Paste the curl command.
6. Click "Continue" and then "Import".

You can create a complete Postman collection by importing each curl command and organizing them into folders.

## Using with Environment Variables in Postman

For ease of use, you can set up environment variables in Postman:

1. Create a new environment (gear icon in the top right).
2. Add variables like:
   - `base_url`: `http://localhost:8082/api/v1`
   - `user_id`: (your test user ID)
   - `profile_id`: (a profile ID for testing)
   - `match_id`: (a match ID for testing)
3. Then replace hardcoded values with variables in your requests:
   - `{{base_url}}/matches/{{match_id}}/messages?user_id={{user_id}}`
