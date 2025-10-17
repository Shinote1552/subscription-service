CREATE TABLE IF NOT EXIST subscriptions(
    id BIGSERIAL PRIMARY KEY,
    service_name TEXT NOT NULL,
    price BIGINT NOT NULL CHECK (price > 0)
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_subscriptions_user_id ON subscriptions USING BTREE (user_id);
CREATE INDEX IF NOT EXISTS idx_subscriptions_service_name ON subscriptions USING BTREE (service_name);
CREATE INDEX IF NOT EXISTS idx_subscriptions_dates ON subscriptions USING BTREE (start_date, end_date);