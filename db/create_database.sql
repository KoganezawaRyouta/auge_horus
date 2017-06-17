-- mysql -u root
-- show databases
-- CREATE DATABASE batch_coincheck
-- ALTER DATABASE batch_coincheck DEFAULT CHARACTER SET=utf8;
-- GRANT ALL ON batch_coincheck.* to coin_app@localhost;
-- SET PASSWORD FOR coin_app@localhost=password('test');

CREATE DATABASE IF NOT EXISTS augehorus_development;
CREATE DATABASE IF NOT EXISTS augehorus_test;

GRANT ALL PRIVILEGES ON augehorus_development.* TO 'coin_app'@'%' IDENTIFIED BY 'test';
GRANT ALL PRIVILEGES ON augehorus_test.* TO 'coin_app'@'%' IDENTIFIED BY 'test';

-- ログインする時
-- mysql -h localhost -u coin_app -p
-- use augehorus_development;


-- integration
-- mysql -u root
-- how databases
-- REATE DATABASE augehorus_integration
-- RLTER DATABASE augehorus_integration DEFAULT CHARACTER SET=utf8;
-- RRANT ALL ON augehorus_integration.* to coin_app@10.0.1.10;
-- RET PASSWORD FOR coin_app@10.0.1.10=password('test');

