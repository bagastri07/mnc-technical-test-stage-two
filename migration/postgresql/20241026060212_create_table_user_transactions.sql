-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_transactions(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    transaction_type VARCHAR NOT NULL,
    action_type VARCHAR NOT NULL,
    amount NUMERIC NOT NULL CHECK (amount > 0),
    balance_before NUMERIC NOT NULL,
    balance_after NUMERIC NOT NULL,
    remarks TEXT,
    status VARCHAR NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_transactions;
-- +goose StatementEnd
