CREATE TABLE "users" (
  "id" bigint,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "username" text NOT NULL,
  "password" text NOT NULL,
  "verified_at" timestamp DEFAULT(NULL),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "email_unique" ON "users" ("email");
CREATE UNIQUE INDEX "username_unique" ON "users" ("username");

-- Trigger for updated_at field
CREATE OR REPLACE FUNCTION updated_timestamp_func()
RETURNS TRIGGER
LANGUAGE plpgsql AS
'
BEGIN
    NEW.updated_at = now();
    NEW.created_at = OLD.created_at;
    RETURN NEW;
END;
';

DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER trigger_update_timestamp
                    BEFORE UPDATE ON %I
                    FOR EACH ROW EXECUTE PROCEDURE updated_timestamp_func()', t,t);
    END loop;
END;
$$ language 'plpgsql';