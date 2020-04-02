CREATE TABLE clients (
  clientid serial primary key,
  firstname text,
  lastname text,
  email text,
  created timestamptz,
  phone text
);

CREATE TABLE addresses (
  addressid serial primary key,
  clientid integer references clients(clientid) not null,
  country text,
  addresssline1 text,
  addresssline2 text,
  city text,
  zip text,
  state text
);

CREATE TABLE promotions (
  promotionid serial primary key,
  code text,
  modifier float(2)
);

CREATE TABLE categories (
  categoryid serial primary key,
  name text
);

CREATE TABLE products (
  productid serial primary key,
  name text,
  description text,
  price int,
  discount float(2),
  categoryid integer references categories(categoryid),
  created timestamptz,
  thumnbnail text,
  public boolean,
  identifier text
);

CREATE TABLE orders (
  orderid serial primary key,
  clientid integer references clients(clientid) not null,
  promotionid integer references promotions(promotionid),
  shippingid integer references adresses(adressid) not null,
  deliveryid integer references adresses(adressid) not null,
  status text,
  created timestamptz,
  token text
);

CREATE TABLE orderitems (
  orderid integer references orders(orderid) not null,
  productid integer references products(productid) not null,
  productid integer references productstock(productstockid) not null,
  quantity integer,
  option text
);

CREATE TABLE productimages (
  productid integer references products(productid) not null,
  image text
);

CREATE TABLE productstock (
  productstockid serial primary key,
  productid integer references products(productid) not null,
  option text,
  quantity integer
);

CREATE TABLE settings (
  settingid serial primary key,
  name text,
  value text
);

INSERT INTO settings VALUES(DEFAULT, 'password', '');
