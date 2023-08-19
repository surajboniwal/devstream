-- Drop foreign keys

-- Drop triggers
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "users";

-- Drop indexes
DROP INDEX IF EXISTS "email_unique";
DROP INDEX IF EXISTS "username_unique";

-- Drop tables
DROP TABLE IF EXISTS "users";
