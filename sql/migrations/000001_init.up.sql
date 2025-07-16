CREATE TABLE orders (
	id         varchar(36) NOT NULL PRIMARY KEY, 
	Price      DECIMAL(10,2) NOT NULL,
	Tax        DECIMAL(10,2) NOT NULL,
	Final_Price DECIMAL(10,2) NOT NULL
);