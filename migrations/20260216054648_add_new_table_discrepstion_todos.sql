-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos ADD COLUMN description TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todos DROP COLUMN IF EXISTS description;
-- +goose StatementEnd
