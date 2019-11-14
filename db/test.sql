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