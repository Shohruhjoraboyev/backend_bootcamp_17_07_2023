create table Products(
    id serial primary key,
    product_name varchar,
    price numeric,
    category_id serial REFERENCES Categories(id)
);

create table Branches(
    id serial primary key,
    branches_name varchar,
    adress varchar 
);

create table Branch_products(
     branch_id serial primary key REFERENCES Branches(id),
     product_id serial primary key REFERENCES Products(product_id),
     quantity integer
    
);

create table BranchPrTransaction(
     id serial primary key,
     branch_id serial primary key REFERENCES Branches(id),
     product_id serial REFERENCES Products(product_id) ,
     transactionType text,
     quantity integer  REFERENCES Branch_products(quantity),
     created_at timestamp
)

create table Categories(
    id serial primary key,
    product_name varchar REFERENCES Products(product_name)
)

