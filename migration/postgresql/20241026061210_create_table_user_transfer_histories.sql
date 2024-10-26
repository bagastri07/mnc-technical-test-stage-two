-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_transfer_histories(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    recipient_transaction_id UUID NOT NULL REFERENCES user_transactions(id),
    sender_transaction_id UUID NOT NULL REFERENCES user_transactions(id),
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_transfer_histories;
-- +goose StatementEnd
