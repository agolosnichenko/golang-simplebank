CREATE TABLE "users" (
    "username" TEXT PRIMARY KEY,
    "hashed_password" TEXT NOT NULL,
    "full_name" TEXT NOT NULL,
    "email" TEXT UNIQUE NOT NULL,
    "password_changed_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username") ON DELETE CASCADE;
-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");