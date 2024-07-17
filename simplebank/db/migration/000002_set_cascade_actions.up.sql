-- Drop existing foreign key constraints
ALTER TABLE "entries" DROP CONSTRAINT IF EXISTS "entries_account_id_fkey";
ALTER TABLE "transfers" DROP CONSTRAINT IF EXISTS "transfers_from_account_id_fkey";
ALTER TABLE "transfers" DROP CONSTRAINT IF EXISTS "transfers_to_account_id_fkey";

-- Add new foreign key constraints with ON DELETE CASCADE
ALTER TABLE "entries" ADD CONSTRAINT "entries_account_id_fkey"
    FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD CONSTRAINT "transfers_from_account_id_fkey"
    FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD CONSTRAINT "transfers_to_account_id_fkey"
    FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
