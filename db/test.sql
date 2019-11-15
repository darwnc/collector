-- 版本 --当前时间
SELECT VERSION(), CURRENT_DATE;
SELECT USER();

SHOW DATABASES;

create database pets;

SHOW DATABASES;

use pets

CREATE TABLE cats
(
  id              INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
  name            VARCHAR(150) NOT NULL,                # Name of the cat
  owner           VARCHAR(150) NOT NULL,                # Owner of the cat
  birth           DATE NOT NULL,                        # Birthday of the cat
  PRIMARY KEY     (id)                                  # Make the id the primary key
);

CREATE TABLE shop (
    article INT UNSIGNED  DEFAULT '0000' NOT NULL,
    dealer  CHAR(20)      DEFAULT ''     NOT NULL,
    price   DECIMAL(16,2) DEFAULT '0.00' NOT NULL,
    PRIMARY KEY(article, dealer));
INSERT INTO shop VALUES
    (1,'A',3.45),(1,'B',3.99),(2,'A',10.99),(3,'B',1.45),
    (3,'C',1.69),(3,'D',1.25),(4,'D',19.95);


SHOW TABLES;

DESCRIBE cats;

INSERT INTO cats ( name, owner, birth) VALUES
  ( 'Sandy', 'Lennon', '2015-01-03' ),
  ( 'Cookie', 'Casey', '2013-11-13' ),
  ( 'Charlie', 'River', '2016-05-21' );
-- 增加gender在name之后
ALTER TABLE cats ADD gender CHAR(1) AFTER name;

SHOW CREATE TABLE cats
-- 删除gender
ALTER TABLE cats DROP gender;

-- and or 操作符
SELECT * FROM `cats` WHERE name='Sandy' or name='Cookie'
-- 排序
SELECT name, birth FROM `cats` ORDER BY birth;
SELECT name, birth FROM `cats` ORDER BY birth DESC;

SELECT name, birth FROM `cats` ORDER BY name, birth DESC;

-- 计算年龄
SELECT name, birth, CURDATE(),
       TIMESTAMPDIFF(YEAR,birth,CURDATE()) AS age
       FROM `cats` ORDER BY age;
SELECT name, birth, MONTH(birth) FROM `cats`;
SELECT name, birth FROM `cats` WHERE MONTH(birth) = 5;
--  YEAR(), MONTH(), and DAYOFMONTH(). MONTH()

SELECT '2018-10-31' + INTERVAL 1 DAY;
-- 2018-11-01 加一天 2018-11-30

-- To test for NULL, use the IS NULL and IS NOT NULL operators, as shown here:
SELECT 1 IS NULL, 1 IS NOT NULL;
SELECT 1 = NULL, 1 <> NULL, 1 < NULL, 1 > NULL;
SELECT 0 IS NULL, 0 IS NOT NULL, '' IS NULL, '' IS NOT NULL;

-- 满足五个字符的 5个_
SELECT * FROM `cats` WHERE name LIKE '_____';

SELECT * FROM `cats` WHERE REGEXP_LIKE(name, '^.{5}$');

SELECT name count(*) FROM `cats` GROUP BY name

-- Find the number, dealer, and price of the most expensive article.
SELECT MAX(article) AS article FROM shop;

SELECT article, dealer, price
FROM   shop
WHERE  price=(SELECT MAX(price) FROM shop);

SELECT s1.article, s1.dealer, s1.price
FROM shop s1
LEFT JOIN shop s2 ON s1.price < s2.price
WHERE s2.article IS NULL;

SELECT article, dealer, price
FROM shop
ORDER BY price DESC
LIMIT 1;
-- Find the highest price per article.
SELECT article, MAX(price) AS price
FROM   shop
GROUP BY article
ORDER BY article;

-- For each article, find the dealer or dealers with the most expensive price.
SELECT article, dealer, price
FROM   shop s1
WHERE  price=(SELECT MAX(s2.price)
              FROM shop s2
              WHERE s1.article = s2.article)
ORDER BY article;

SELECT s1.article, dealer, s1.price
FROM shop s1
JOIN (
  SELECT article, MAX(price) AS price
  FROM shop
  GROUP BY article) AS s2
  ON s1.article = s2.article AND s1.price = s2.price
ORDER BY article;

SELECT s1.article, s1.dealer, s1.price
FROM shop s1
LEFT JOIN shop s2 ON s1.article = s2.article AND s1.price < s2.price
WHERE s2.article IS NULL
ORDER BY s1.article;
-- window function
WITH s1 AS (
   SELECT article, dealer, price,
          RANK() OVER (PARTITION BY article
                           ORDER BY price DESC
                      ) AS `Rank`
     FROM shop
)
SELECT article, dealer, price
  FROM s1
  WHERE `Rank` = 1
ORDER BY article;
-- find the articles with the highest and lowest price you can do this:
SELECT @min_price:=MIN(price),@max_price:=MAX(price) FROM shop;
SELECT * FROM shop WHERE price=@min_price OR price=@max_price;


-- Using Foreign Keys --------

CREATE TABLE person (
    id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name CHAR(60) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE shirt (
    id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
    style ENUM('t-shirt', 'polo', 'dress') NOT NULL,
    color ENUM('red', 'blue', 'orange', 'white', 'black') NOT NULL,
    owner SMALLINT UNSIGNED NOT NULL REFERENCES person(id),
    PRIMARY KEY (id)
);
INSERT INTO person VALUES (NULL, 'Antonio Paz');
SELECT @last := LAST_INSERT_ID();

INSERT INTO shirt VALUES
(NULL, 'polo', 'blue', @last),
(NULL, 'dress', 'white', @last),
(NULL, 't-shirt', 'blue', @last);

INSERT INTO person VALUES (NULL, 'Lilliana Angelovska');

SELECT @last := LAST_INSERT_ID();

INSERT INTO shirt VALUES
(NULL, 'dress', 'orange', @last),
(NULL, 'polo', 'red', @last),
(NULL, 'dress', 'blue', @last),
(NULL, 't-shirt', 'white', @last);
SELECT * FROM person;
SELECT * FROM shirt;

SELECT s.* FROM person p INNER JOIN shirt s
   ON s.owner = p.id
 WHERE p.name LIKE 'Lilliana%'
   AND s.color <> 'white';

CREATE TABLE animals (
    grp ENUM('fish','mammal','bird') NOT NULL,
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    name CHAR(30) NOT NULL,
    PRIMARY KEY (grp,id)
) ENGINE=MyISAM;

INSERT INTO animals (grp,name) VALUES
    ('mammal','dog'),('mammal','cat'),
    ('bird','penguin'),('fish','lax'),('mammal','whale'),
    ('bird','ostrich');
insert into animals (grp,name) values ('mammal','dont know');
SELECT * FROM animals ORDER BY grp,id;
-----------------------------------------------------------
select DISTINCT name from `y_user`;
SELECT * from `y_user` WHERE id<>1;
select name from `y_user`;
-- 开启事务
begin;
insert into 
commit;
-- 提交事务
rollback;
-- 回滚事务

-- ALTER TABLE tabel  DROP i; 
--  删除字段（列）如果数据表中只剩余一个字段则无法使用DROP来删除字段
--  ALTER TABLE tabel ADD i INT;
-- 执行以上命令后，i 字段会自动添加到数据表字段的末尾。
-- ALTER TABLE tabel ADD i INT FIRST; 指定字段的位置
-- ALTER TABLE tabel ADD i INT AFTER c;在c字段之后

-- 把字段 c 的类型从 CHAR(1) 改为 CHAR(10)，可以执行以下命令
-- ALTER TABLE tabel MODIFY c CHAR(10);

-- 尝试以下实例将数据表 testalter_tbl 重命名为 alter_tbl：
-- ALTER TABLE testalter_tbl RENAME TO alter_tbl;