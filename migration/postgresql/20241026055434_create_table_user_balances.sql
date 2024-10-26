-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_balances(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    balance NUMERIC NOT NULL CHECK (balance >= 0),
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_balances;
-- +goose StatementEnd
