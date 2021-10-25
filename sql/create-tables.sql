CREATE TABLE employees (
Id varchar(255) NOT NULL,
FirstName varchar(255) not null,
MiddleName varchar(255) NOT NULL,
LastName varchar(255) NOT NULL,
Gender varchar(255),
Salary float not null,
DOB DATE not null,
Email varchar(255),
Phone int,
AddressLine1 varchar(255) NOT NULL,
AddressLine2 varchar(255),
State varchar(255) NOT NULL,
PostCode int NOT NULL,
TFN int NOT NULL,
SuperBalance float NOT NULL,
PRIMARY KEY (Id)
);


insert into employees values (1,'Vinodh','K','L','Male',555.55,'1993-12-10','vinod@gmail.com',
4634645,'Lonsdale','street','vic',3000,1354354,4645.0);



insert into employees values (2,'Santos','K','L','Male',444.44,'1994-12-10','santos@gmail.com',
0354345,'Spencer','street','vic',3000,346365,4645.0);