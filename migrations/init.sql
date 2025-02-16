
CREATE TABLE IF NOT EXISTS shop_stock (
    product_name TEXT NOT NULL UNIQUE,
    price INT NOT NULL
);

INSERT INTO shop_stock (product_name, price) VALUES
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500)
ON CONFLICT (product_name) DO NOTHING;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    balance INT NOT NULL DEFAULT 1000,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS purchases (
    transaction_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    merch_name TEXT NOT NULL,
    price INT NOT NULL,
    quantity INT NOT NULL,
    total INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) references users (id),
    FOREIGN KEY (merch_name) references shop_stock (product_name)
);

CREATE TABLE IF NOT EXISTS coins_transfers (
    transfer_id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    amount INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (sender_id) references users (id),
    FOREIGN KEY (receiver_id) references users (id)
);


