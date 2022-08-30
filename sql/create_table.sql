CREATE TABLE products (
    id UUID not null primary key,
    name varchar(20) not null,
    quantity int not null,
    price int not null,
    original_price int not null
);
CREATE TABLE basket (
    id UUID not null primary key,
    createtime timestamp not null
);
CREATE TABLE  chek (
  product_id UUID not null REFERENCES products(id),
  basket_id UUID not null REFERENCES basket(id),
  quantity int not null
);
CREATE TABLE store (
    budget int not null ,
    profit int not null
);
