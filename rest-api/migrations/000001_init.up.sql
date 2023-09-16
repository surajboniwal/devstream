CREATE TABLE "users" (
  "id" bigint PRIMARY KEY NOT NULL,
  "name" text NOT NULL,
  "username" text NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "verified_at" timestamp DEFAULT null,
  "updated_at" timestamp DEFAULT null,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "stream_keys" (
  "id" bigint PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "name" text NOT NULL,
  "key" bigint NOT NULL,
  "updated_at" timestamp DEFAULT null,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "videos" (
  "id" bigint PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "title" text NOT NULL,
  "description" text NOT NULL DEFAULT '',
  "thumbnail_id" bigint NOT NULL,
  "updated_at" timestamp DEFAULT null,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "images" (
  "id" bigint PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "updated_at" timestamp DEFAULT null,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX "email_unique" ON "users" ("email");
CREATE UNIQUE INDEX "username_unique" ON "users" ("username");
CREATE INDEX "streamkey_id_user_id" ON "stream_keys" ("id", "user_id");
CREATE UNIQUE INDEX "streamkey_key_unique" ON "stream_keys" ("user_id", "key");
CREATE UNIQUE INDEX "streamkey_name_unique" ON "stream_keys" ("user_id", "name");

ALTER TABLE "stream_keys" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "images" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "videos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "videos" ADD FOREIGN KEY ("thumbnail_id") REFERENCES "images" ("id");

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