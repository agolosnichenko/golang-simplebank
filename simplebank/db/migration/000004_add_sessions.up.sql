CREATE TABLE "sessions" (
    "id" uuid PRIMARY KEY,
    "username" TEXT NOT NULL,
    "refresh_token" TEXT NOT NULL,
    "user_agent" TEXT NOT NULL,
    "client_ip" TEXT NOT NULL,
    "is_blocked" BOOLEAN NOT NULL DEFAULT FALSE,
    "expires_at" TIMESTAMPTZ NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username") ON DELETE CASCADE;