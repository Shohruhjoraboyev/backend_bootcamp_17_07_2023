/*HOMEWORK 1.*/

Table yaratish 1-vazifa
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL UNIQUE,
    description TEXT,
    price NUMERIC,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

/*2-vazifa*/

ALTER TABLE products
ADD CONSTRAINT unique_name UNIQUE (name);

ALTER TABLE products
ALTER COLUMN name SET NOT NULL;

/*3-vazifa*/

ALTER TABLE products
DROP CONSTRAINT unique_name;

/*4-vazifa*/

SELECT *
FROM products
WHERE name LIKE '%a%' AND price > 5000;

/*5-vazifa*/

SELECT name, created_at
FROM products
WHERE created_at BETWEEN NOW() - INTERVAL '5 days' AND NOW()
ORDER BY created_at DESC;

/*6-vazifa*/

SELECT name, created_at, updated_at
FROM products
WHERE (name LIKE '%burger%' OR description LIKE '%burger%')
      AND (created_at = updated_at OR updated_at IS NULL);
