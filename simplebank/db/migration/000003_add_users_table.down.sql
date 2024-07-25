-- Reverse the addition of the unique constraint on the "accounts" table
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

-- Reverse the addition of the foreign key on the "accounts" table
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

-- Drop the "users" table
DROP TABLE IF EXISTS "users";
