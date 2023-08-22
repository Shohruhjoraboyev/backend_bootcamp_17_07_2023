
-- 2. products table yaratish:id primary key bo'lishi kerak

CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  description TEXT,
  price NUMERIC(10,2),
  createdAt TIMESTAMP,
  updatedAt TIMESTAMP
);

--3. 1.name columnga unique va not null qo'shish

ALTER TABLE products
ALTER COLUMN name SET NOT NULL,
ADD CONSTRAINT unique_name UNIQUE (name);

--3. 2. uniqueni olib tashlash
ALTER TABLE products
DROP CONSTRAINT IF EXISTS unique_name;


--4. nameda "a" harfi qatnashgan va narxi 5000dan katta bo'lgan productlarni select qilish

INSERT INTO products (name, description, price, createdAt, updatedAt)
VALUES
  ('Olma', 'Qizil olma', 3500, now(), now()),
  ('Anor', 'Katta shirin anor', 3200, now(), now()),
  ('Shaftoli', ' Yangi Tuksiz shaftoli', 4000, now(), now()),
  ('O''rik', 'qandek shirn o''rik', 2800, now(), now()),
  ('Banana', ' Yangi Suvli banan', 3000, now(), now()),
  ('Nok', 'Toza noklar', 4500, now(), now()),
  ('Qulupnay', ' Yangi Shirin qulupnay', 3100, now(), now()),
  ('Kivi', 'export kivi', 3800, now(), now()),
  ('Ananas', 'Yangi ananas', 5000, now(), now()),
  ('Mango', 'Mango mevasi', 4200, now(), now()
);
----4. nameda "a" harfi qatnashgan va narxi 5000dan katta bo'lgan productlarni select qilish
SELECT * FROM products
WHERE name LIKE '%a%' AND price > 5000;


--5.  oxirgi 5kunda kiritilgan productlarni faqat name,created_atlarini, kiritilgan vaqtini kamayish tartibida chiqarish(between ishlatilsin)

SELECT name, createdat
FROM products
WHERE createdat BETWEEN current_date - interval '5 days' AND current_date
ORDER BY createdat ASC;

--Update updateAt

UPDATE Products
SET updatedAt = current_timestamp
WHERE id IN (1, 3, 5, 7);

--6. name yoki descriptionda "yangi" so'zi qatnashgan va yaratilgandan keyin update qilingan productlarni chiqarish(created_at yaratilgan vaqt, updated_at update qilingan vaqt. update qilinmagan productda updated_at created_atga teng bo'ladi)

SELECT *FROM Products
WHERE (name ILIKE '%yangi%' OR description ILIKE '%yangi%')
  AND updated_at > created_at;