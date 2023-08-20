-- Drop foreign keys

-- Drop triggers
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "users";
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "stream_keys";

-- Drop indexes
DROP INDEX IF EXISTS "email_unique";
DROP INDEX IF EXISTS "username_unique";
DROP INDEX IF EXISTS "user_key_unique";

-- Drop tables
DROP TABLE IF EXISTS "stream_keys";
DROP TABLE IF EXISTS "users";
