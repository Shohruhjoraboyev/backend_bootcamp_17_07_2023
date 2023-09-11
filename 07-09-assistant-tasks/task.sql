create table company(
    id int PRIMARY KEY,
    name varchar
);

insert into company(id,name) values 
(1,'apple'),
(2,'samsung'),
(3,'xiaomi');


create table sale (
    id serial,
    product_name varchar,
    count int, 
    company_id  int references company(id)
);

insert into sale (product_name, count, company_id) values 
('Iphone 14', 50, 1),
('Iphone 13', 20, 1),
('Iphone 12', 10, 1),
('Samsung Galaxy S20', 50, 2),
('Samsung Galaxy S21', 30, 2),
('Xiaomi Redmi K50', 10, 3),
('Xiaomi Mi 10 Lite', 20, 3),
('Xiaomi Redmi 13 Lite', 10, 3);

-----------------------------------------------------
-- eng kop product sotgan comaniyalar chiqarish kerak
--  (nechta bolasa hammasi
--   Masalan :
--             Apple    80
--             Samsung 80
-- )

-------------------------
WITH CompanyTotal AS (
    SELECT c.name AS name, SUM(s.count) AS count
    FROM company c
    JOIN sale s ON c.id = s.company_id
    GROUP BY c.name
)
SELECT name, count
FROM CompanyTotal
WHERE count = (SELECT MAX(count) FROM CompanyTotal);
