DROP TABLE IF EXISTS positions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS portfolios;
DROP TABLE IF EXISTS money;

CREATE TABLE users
(
    id    SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL
);
CREATE TABLE portfolios
(
    id      SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    name    VARCHAR(255) NOT NULL
);
CREATE TABLE money
(
    id           SERIAL PRIMARY KEY,
    portfolio_id INTEGER REFERENCES portfolios (id) ON DELETE CASCADE,
    amount       DOUBLE PRECISION NOT NULL
);
CREATE TABLE positions
(
    id           SERIAL PRIMARY KEY,
    portfolio_id INTEGER REFERENCES portfolios (id) ON DELETE CASCADE,
    ticker       VARCHAR(255),
    price        DOUBLE PRECISION,
    quantity     DOUBLE PRECISION,
    close_price  DOUBLE PRECISION,
    opened_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    closed_at    TIMESTAMP,
    profit       DOUBLE PRECISION
);

INSERT INTO users (email)
VALUES ('user@example.com')
RETURNING id;

INSERT INTO portfolios (user_id, name)
VALUES (1, 'Portfolio 1'),
       (1, 'Portfolio 2');

INSERT INTO money (portfolio_id, amount)
VALUES (1, 10000.0),
       (2, 5000.0);

INSERT INTO positions (portfolio_id, ticker, price, quantity)
VALUES (1, 'EDV', 1892.0, 10.0),
       (1, 'SDR', 440.55, 10.0),
       (1, 'ABDN', 222.5, -10.0),
       (2, 'AAPL', 142.0, 5.0),
       (2, 'GOOGL', 2750.0, 2.0);