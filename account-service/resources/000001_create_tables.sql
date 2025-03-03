CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY ,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    account_type VARCHAR(20) NOT NULL,
    balance DECIMAL(15,2) NOT NULL DEFAULT 0.00,
    status VARCHAR(10) NOT NULL DEFAULT 'active',
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

--transactions.sql
create table IF NOT EXISTS transactions (
    id UUID primary key,
    account UUID not null,
    amount decimal(15,2) not null,
    type varchar(10) not null,
    timestamp timestamp not null,
    foreign key (account) references accounts(id)
);

--users.sql
create table IF NOT EXISTS users (
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email varchar(100) unique not null,
    type varchar(20) not null,
    password varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);  

-- Create index if not exists
CREATE INDEX IF NOT EXISTS idx_accounts_email ON accounts(email);

-- Add constraints if they don't exist
DO $$ 
BEGIN 
    -- Check and add balance_non_negative constraint
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'balance_non_negative'
    ) THEN
        ALTER TABLE accounts ADD CONSTRAINT balance_non_negative CHECK (balance >= 0);
    END IF;

    -- Check and add valid_status constraint
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'valid_status'
    ) THEN
        ALTER TABLE accounts ADD CONSTRAINT valid_status CHECK (status IN ('active', 'inactive', 'frozen', 'closed'));
    END IF;

    -- Check and add valid_account_type constraint
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'valid_account_type'
    ) THEN
        ALTER TABLE accounts ADD CONSTRAINT valid_account_type CHECK (account_type IN ('savings', 'checking', 'credit'));
    END IF;
END $$;

-- Add trigger function if it doesn't exist
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Drop trigger if exists and create new one
DROP TRIGGER IF EXISTS update_accounts_updated_at ON accounts;
CREATE TRIGGER update_accounts_updated_at
    BEFORE UPDATE ON accounts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();