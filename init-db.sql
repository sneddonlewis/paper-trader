DROP TABLE IF EXISTS positions;

CREATE TABLE positions
(
    id        SERIAL PRIMARY KEY,
    ticker    VARCHAR(255),
    price     DOUBLE PRECISION,
    quantity  DOUBLE PRECISION
);

INSERT INTO positions (ticker, price, quantity)
VALUES ('EDV', 1892.0, 10.0),
       ('SDR', 440.55, 10.0),
       ('ABDN', 222.5, -10.0);