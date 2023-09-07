/* TASK
4. array_agg() ishlatib branchda qanaqa categoriyadagi productlar borligini chiqaring
Example:

 branch      |  categories
Chilonzor   | {meva, ichimlik, kiyim}
MGorkiy     | {meva}  


*/

CREATE TABLE Branch (
  id SERIAL PRIMARY KEY,
  name Varchar(255),
  category_id INT
);

INSERT INTO Branch (name, category_id)
(VALUES
    ('Beruniy',Array[1,2,4]),
    ('Maksim go''rkiy',Array[2,3,5]),
    ('Chilonzor',Array[4,9]),
    ('Pushkin',Array[6,7,8])
);

-- Create the "categories" table
CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  Category_name TEXT
);

-- Insert randomly generated data into the "categories" table
INSERT INTO categories (Category_name)
VALUES
  ('Meva'),
  ('Ichimlik'),
  ('Sabzavot'),
  ('Texnika'),
  ('Shirinlik'),
  ('Noutbook'),
  ('Planshet'),
  ('Mobil_Telefon');

 --Doing with ARRAY_AGG

 SELECT b.name AS branch, array_agg(c.Category_name) AS categories
FROM Branch AS b
JOIN categories AS c ON b.category_id = c.id
GROUP BY b.name;

--Output
/*
branch        |           categories
-----------------+---------------------------------
 Beruniy         | {Meva,Ichimlik,Shirinlik}
 Chilonzor       | {Shirinlik,Noutbook}
 Maksim go'rkiy  | {Ichimlik,Sabzavot,Noutbook}
 Pushkin         | {Planshet,Mobil_Telefon}

*/
