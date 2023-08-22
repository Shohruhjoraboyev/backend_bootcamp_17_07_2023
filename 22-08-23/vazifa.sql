-- task 2
create table products(
    id serial primary key,
    name varchar(100),
    description text,
    price int,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

-- task 3/1
alter table products 
add constraint name_unique unique(name), 
alter column name set not null;

-- task 3/2
alter table products
drop constraint name_unique;


-- task 4
select * from products 
where name ilike '%a%' and price>5000;

-- task 5   
select name,created_at
from products
where created_at between now() - interval '5 days' and now() 
order by created_at desc;

-- task 6
select * from products
where (name ilike '%burger%' or description ilike '%burger%')
and updated_at > created_at;


-- insert into products(name,description,price)
-- values 
-- ('olma','yashil',12000),
-- ('guruch','alanga',15500),
-- ('yog''','1L',14800),
-- ('qaymoq','400gr',18000),
-- ('asal','yantoq',65000),
-- ('go''sht','mol',85000),
-- ('piyoz','oq',5600);


