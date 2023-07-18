CREATE TABLE positions (
    ticker VARCHAR(255),
    direction VARCHAR(255),
    price FLOAT,
    quantity FLOAT
);
insert into positions values (
    'EDV', 'long', 1892.0, 10.0
);

CREATE TABLE technologies (
  name    VARCHAR(255),
  details VARCHAR(255)
);
insert into technologies values (
  'Go', 'An open source programming language that makes it easy to build simple and efficient software.'
);
insert into technologies values (
  'JavaScript', 'A lightweight, interpreted, or just-in-time compiled programming language with first-class functions.'
);
insert into technologies values (
  'MySQL', 'A powerful, open source object-relational database'
);
